package vimg

// Compiler version constants.
const (
	SemverMajor    = 3
	SemverMinor    = 0
	SemverRevision = 0
)

var (
	// Bootloader contains all of the bytes that will be written as the image's bootloader.
	Bootloader = []byte{
		0xFA, 0x31, 0xC0, 0x8E, 0xD8, 0x8E, 0xD0, 0xBC, 0x00, 0x7C, 0xB8, 0x00, 0x10, 0x8E, 0xC0, 0xFB, 0x66, 0xB8, 0x01, 0x00, 0x00, 0x00, 0x31, 0xDB, 0xB9, 0x00, 0x03, 0xE8, 0xDA, 0x00, 0x66, 0xB8, 0x43, 0x00, 0x00, 0x00, 0x66, 0xA3, 0x92, 0x7D, 0x66, 0xB8, 0x01, 0x00, 0x00, 0x00, 0x31, 0xDB, 0xB9, 0x00, 0x10, 0xE8, 0xC2,
		0x00, 0x66, 0x31, 0xC0, 0x26, 0xA0, 0xF1, 0x01, 0x83, 0xF8, 0x00, 0x75, 0x03, 0xB8, 0x04, 0x00, 0xBB, 0x00, 0x02, 0xB9, 0x00, 0x10, 0xE8, 0xAA, 0x00, 0x26, 0xC6, 0x06, 0x10, 0x02, 0xE1, 0x26, 0x80, 0x0E, 0x11, 0x02, 0x80, 0x26, 0xC7, 0x06, 0x24, 0x02, 0x00, 0xDE, 0x26, 0xC6, 0x06, 0x27, 0x02, 0x01, 0x26, 0x66, 0xC7,
		0x06, 0x28, 0x02, 0x00, 0xE0, 0x01, 0x00, 0xFC, 0xBE, 0x00, 0x31, 0xBF, 0x00, 0xE0, 0x8B, 0x0E, 0x20, 0x30, 0xF3, 0xA4, 0x26, 0x66, 0x8B, 0x16, 0xF4, 0x01, 0x66, 0xC1, 0xE2, 0x04, 0x66, 0x83, 0xFA, 0x00, 0x74, 0x54, 0x66, 0x81, 0xFA, 0x00, 0xFE, 0x00, 0x00, 0x72, 0x28, 0x66, 0xB8, 0x7F, 0x00, 0x00, 0x00, 0x31, 0xDB,
		0xB9, 0x00, 0x20, 0xE8, 0x53, 0x00, 0xB9, 0x00, 0x7F, 0xE8, 0x77, 0x00, 0x66, 0x81, 0xEA, 0x00, 0xFE, 0x00, 0x00, 0x81, 0x06, 0x6C, 0x7D, 0x00, 0xFE, 0x80, 0x16, 0x6E, 0x7D, 0x00, 0xEB, 0xC9, 0x66, 0x89, 0xD0, 0x66, 0xC1, 0xE8, 0x09, 0x66, 0xF7, 0xC2, 0xFF, 0x01, 0x00, 0x00, 0x74, 0x02, 0x66, 0x40, 0x31, 0xDB, 0xB9,
		0x00, 0x20, 0xE8, 0x1F, 0x00, 0x66, 0x89, 0xD1, 0x66, 0xD1, 0xE9, 0xE8, 0x40, 0x00, 0xFA, 0xB8, 0x00, 0x10, 0x8E, 0xD8, 0x8E, 0xC0, 0x8E, 0xE0, 0x8E, 0xE8, 0x8E, 0xD0, 0xBC, 0x00, 0xE0, 0xEA, 0x00, 0x00, 0x20, 0x10, 0x66, 0x52, 0xA3, 0x84, 0x7D, 0x89, 0x1E, 0x86, 0x7D, 0x89, 0x0E, 0x88, 0x7D, 0x66, 0x8B, 0x16, 0x92,
		0x7D, 0x66, 0x89, 0x16, 0x8A, 0x7D, 0x66, 0x01, 0x06, 0x92, 0x7D, 0xB4, 0x42, 0xBE, 0x82, 0x7D, 0xB2, 0x80, 0xCD, 0x13, 0x72, 0x17, 0x66, 0x5A, 0xC3, 0x66, 0x52, 0x06, 0x31, 0xC0, 0x8E, 0xC0, 0xB4, 0x87, 0xBE, 0x52, 0x7D, 0xCD, 0x15, 0x72, 0x04, 0x07, 0x66, 0x5A, 0xC3, 0xBE, 0x96, 0x7D, 0xAC, 0x20, 0xC0, 0x74, 0x09,
		0xB4, 0x0E, 0xBB, 0x07, 0x00, 0xCD, 0x10, 0xEB, 0xF2, 0x31, 0xC0, 0xCD, 0x16, 0xCD, 0x19, 0xEA, 0xF0, 0xFF, 0x00, 0xF0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0x02, 0x93, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0x10, 0x93, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x65, 0x72, 0x72, 0x00}
)
