
package main

import (
	"fmt"
)

func RunBytes(bytes []uint16) {
	// Define Registers & Ram
	var reg [4]uint16
	var ram [65536]uint16

	// Load bytes into RAM, starting at 0
	for i := 0; i < len(bytes); i++ {
		ram[i] = bytes[i]
	}

	// Run through the ram
	for head := 0; head < len(ram); head++ {
		inst := ram[head]

		// Check instruction
		switch (inst) {
		case 1: // LOAD_A
			head += 1
			reg[0] = ram[head]
		case 2: // LOAD_B
			head += 1
			reg[1] = ram[head]
		case 3: // LOAD_C
			head += 1
			reg[2] = ram[head]
		case 4: // LOAD_D
			head += 1
			reg[3] = ram[head]

		case 5: // LOAD_A_RAM
			head += 1
			address := ram[head]
			reg[0] = ram[address]
		case 6: // LOAD_B_RAM
			head += 1
			address := ram[head]
			reg[1] = ram[address]
		case 7: // LOAD_C_RAM
			head += 1
			address := ram[head]
			reg[2] = ram[address]
		case 8: // LOAD_D_RAM
			head += 1
			address := ram[head]
			reg[3] = ram[address]

		case 9: // STORE_A
			head += 1
			address := ram[head]
			ram[address] = reg[0]
		case 10: // STORE_B
			head += 1
			address := ram[head]
			ram[address] = reg[1]
		case 11: // STORE_C
			head += 1
			address := ram[head]
			ram[address] = reg[2]
		case 12: // STORE_D
			head += 1
			address := ram[head]
			ram[address] = reg[3]

		case 13: // ADD
			head += 1
			reg1 := ram[head]
			head += 1
			reg2 := ram[head]
			reg[reg1] = reg[reg1] + reg[reg2]
		case 14: // SUB
			head += 1
			reg1 := ram[head]
			head += 1
			reg2 := ram[head]
			reg[reg1] = reg[reg1] - reg[reg2]
		case 15: // MUL
			head += 1
			reg1 := ram[head]
			head += 1
			reg2 := ram[head]
			reg[reg1] = reg[reg1] * reg[reg2]
		case 16: // DIV
			head += 1
			reg1 := ram[head]
			head += 1
			reg2 := ram[head]
			reg[reg1] = reg[reg1] / reg[reg2]

		case 17: // LOG_A
			fmt.Println(reg[0])
		case 18: // LOG_B
			fmt.Println(reg[1])
		case 19: // LOG_C
			fmt.Println(reg[2])
		case 20: // LOG_D
			fmt.Println(reg[3])

		case 21: // LOG_A_CHAR
			fmt.Print(string(reg[0]))
		case 22: // LOG_B_CHAR
			fmt.Print(string(reg[1]))
		case 32: // LOG_C_CHAR
			fmt.Print(string(reg[2]))
		case 33: // LOG_D_CHAR
			fmt.Print(string(reg[3]))

		case 25: // JUMP
			head += 1
			op := ram[head]
			head += 1
			reg1 := ram[head]
			val1 := reg[reg1]
			head += 1
			reg2 := ram[head]
			val2 := reg[reg2]
			head += 1
			loc := ram[head]

			toJump := false
			switch (op) {
				case 0: if val1 == val2 { toJump = true }
				case 1: if val1 != val2 { toJump = true }
				case 2: if val1 < val2  { toJump = true }
				case 3: if val1 > val2  { toJump = true }
				case 4: if val1 <= val2 { toJump = true }
				case 5: if val1 >= val2 { toJump = true }
			}

			if toJump {
				// Jump to <loc>
				head = int(loc)
			}
		}
	}
}

