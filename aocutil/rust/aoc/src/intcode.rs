use std::fs::File;

pub fn prog_stack_reader(file: File) -> Vec<isize> {
    let reading_handler = |line: String, mut col: Vec<isize>| -> Vec<isize> {
        col.extend(line.split(',')
            .map(|s| s.parse::<isize>().unwrap()));
        return col
    };
    crate::io::read_puzzle_file(file, reading_handler)
}

/// Interpret intcode commands at instruction pointer address of program. returns new instruction pointer.
/// instruction pointer will be -1 if program exited on intcode 99.
pub fn intcode_interpreter(prog_stack: &mut Vec<isize>, inst_pointer: usize) -> isize {
    if prog_stack[inst_pointer] == 1 {
        intcode_add(prog_stack, inst_pointer);
        return inst_pointer as isize + 4
    } else if prog_stack[inst_pointer] == 2 {
        intcode_mul(prog_stack, inst_pointer);
        return inst_pointer as isize + 4
    } else if prog_stack[inst_pointer] == 99 {
        return -1
    } else {
        panic!("Unexpected instruction ({}) at adress: {}", prog_stack[inst_pointer], inst_pointer);
    }
}

fn intcode_add(prog_stack: &mut Vec<isize>, inst_pointer: usize) {
    let result = prog_stack[prog_stack[inst_pointer + 1] as usize] + prog_stack[prog_stack[inst_pointer + 2] as usize];
    let target_addr = prog_stack[inst_pointer + 3] as usize;
    prog_stack[target_addr] = result;
}

fn intcode_mul(prog_stack: &mut Vec<isize>, inst_pointer: usize) {
    let result = prog_stack[prog_stack[inst_pointer + 1] as usize] * prog_stack[prog_stack[inst_pointer + 2] as usize];
    let target_addr = prog_stack[inst_pointer + 3] as usize;
    prog_stack[target_addr] = result;
}
