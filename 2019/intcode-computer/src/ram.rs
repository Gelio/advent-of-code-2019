pub struct RAM {
    memory: Vec<isize>,
}

impl RAM {
    pub fn new(memory: Vec<isize>) -> Self {
        Self { memory }
    }

    pub fn get(&mut self, addr: usize) -> isize {
        match self.memory.get(addr) {
            Some(x) => *x,
            None => {
                self.memory.resize(addr + 1, 0);

                0
            }
        }
    }

    pub fn set(&mut self, addr: usize, val: isize) {
        if self.memory.len() <= addr {
            self.memory.resize(addr + 1, 0);
        }

        self.memory[addr] = val;
    }
}
