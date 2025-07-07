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
pub fn intcode_interpreter_v1(prog_stack: &mut Vec<isize>, inst_pointer: usize) -> isize {
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

/// Interpret intcode commands at instruction pointer address of program. returns new instruction pointer.
/// instruction pointer will be -1 if program exited on intcode 99.
pub fn intcode_interpreter_v2(prog_stack: &mut Vec<isize>, inst_pointer: usize, arg: &str) -> isize {
    let inst = prog_stack[inst_pointer] % 100;
    let modes = prog_stack[inst_pointer] << 2;

    if inst[0] == 1 {
        let param_count: isize = 2;

        // add
        return inst_pointer as isize + param_count + 2
    } else if inst[0] == 2 {
        // mul
        return inst_pointer as isize + 4
    } else if inst[0] == 3 {
        // save xxxx3 [Address]
        return inst_pointer as isize + 2;
    } else if inst[0] == 99 {
        return -1
    } else {
        panic!("Unexpected instruction ({}) at adress: {}", prog_stack[inst_pointer], inst_pointer);
    }
}

// 1 1 1 00
fn intcode_add(prog_stack: &mut Vec<isize>, inst_pointer: usize, mut modes: isize) {
    let a = if modes % 1000 > 0 { prog_stack[inst_pointer + 1] } else { prog_stack[prog_stack[inst_pointer + 1] as usize] };
    let b = if modes % 10000 > 0 { prog_stack[inst_pointer + 2] } else { prog_stack[prog_stack[inst_pointer + 2] as usize] };

    let result = a + b;
    let target_addr = prog_stack[inst_pointer + 3] as usize;
    prog_stack[target_addr] = result;
}

fn intcode_mul(prog_stack: &mut Vec<isize>, inst_pointer: usize) {
    let result = prog_stack[prog_stack[inst_pointer + 1] as usize] * prog_stack[prog_stack[inst_pointer + 2] as usize];
    let target_addr = prog_stack[inst_pointer + 3] as usize;
    prog_stack[target_addr] = result;
}
