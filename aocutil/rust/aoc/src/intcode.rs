/// Interpret intcode commands at address of program. returns new address.
/// Address will be -1 if program exited on intcode 99.
pub fn intcode_interpreter(prog_stack: &mut Vec<isize>, addr: usize) -> isize {
    if prog_stack[addr] == 1 {
        intcode_add(prog_stack, addr);
        return addr as isize + 4
    } else if prog_stack[addr] == 2 {
        intcode_mul(prog_stack, addr);
        return addr as isize + 4
    } else if prog_stack[addr] == 99 {
        return -1
    } else {
        panic!("Unexpected instruction ({}) at adress: {}", prog_stack[addr], addr);
    }
}

fn intcode_add(prog_stack: &mut Vec<isize>, addr: usize) {
    let result = prog_stack[prog_stack[addr + 1] as usize] + prog_stack[prog_stack[addr + 2] as usize];
    let target_addr = prog_stack[addr + 3] as usize;
    prog_stack[target_addr] = result;
}

fn intcode_mul(prog_stack: &mut Vec<isize>, addr: usize) {
    let result = prog_stack[prog_stack[addr + 1] as usize] * prog_stack[prog_stack[addr + 2] as usize];
    let target_addr = prog_stack[addr + 3] as usize;
    prog_stack[target_addr] = result;
}
