package unicodeh

import "unicode"

// Unicode property "Uppercase" (known as "Upper", "Uppercase").
// Kind of property: "Binary".
// Based on file "DerivedCoreProperties.txt".
var (
	UppercaseNo  = uppercaseNo  // Value "No" (known as "N", "No", "F", "False").
	UppercaseYes = uppercaseYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	uppercaseNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x40, 0x1}, {0x5b, 0xbf, 0x1}, {0xd7, 0xdf, 0x8}, {0xe0, 0xff, 0x1}, {0x101, 0x137, 0x2}, {0x138, 0x148, 0x2}, {0x149, 0x177, 0x2}, {0x17a, 0x17e, 0x2}, {0x17f, 0x180, 0x1}, {0x183, 0x185, 0x2}, {0x188, 0x18c, 0x4}, {0x18d, 0x192, 0x5}, {0x195, 0x199, 0x4}, {0x19a, 0x19b, 0x1}, {0x19e, 0x1a1, 0x3}, {0x1a3, 0x1a5, 0x2}, {0x1a8, 0x1aa, 0x2}, {0x1ab, 0x1ad, 0x2}, {0x1b0, 0x1b4, 0x4}, {0x1b6, 0x1b9, 0x3}, {0x1ba, 0x1bb, 0x1}, {0x1bd, 0x1c3, 0x1}, {0x1c5, 0x1c6, 0x1}, {0x1c8, 0x1c9, 0x1}, {0x1cb, 0x1cc, 0x1}, {0x1ce, 0x1dc, 0x2}, {0x1dd, 0x1ef, 0x2}, {0x1f0, 0x1f2, 0x2}, {0x1f3, 0x1f5, 0x2}, {0x1f9, 0x233, 0x2}, {0x234, 0x239, 0x1}, {0x23c, 0x23f, 0x3}, {0x240, 0x242, 0x2}, {0x247, 0x24f, 0x2}, {0x250, 0x36f, 0x1}, {0x371, 0x373, 0x2}, {0x374, 0x375, 0x1}, {0x377, 0x37e, 0x1}, {0x380, 0x385, 0x1}, {0x387, 0x38b, 0x4}, {0x38d, 0x390, 0x3}, {0x3a2, 0x3ac, 0xa}, {0x3ad, 0x3ce, 0x1}, {0x3d0, 0x3d1, 0x1}, {0x3d5, 0x3d7, 0x1}, {0x3d9, 0x3ef, 0x2}, {0x3f0, 0x3f3, 0x1}, {0x3f5, 0x3f6, 0x1}, {0x3f8, 0x3fb, 0x3}, {0x3fc, 0x430, 0x34}, {0x431, 0x45f, 0x1}, {0x461, 0x481, 0x2}, {0x482, 0x489, 0x1}, {0x48b, 0x4bf, 0x2}, {0x4c2, 0x4ce, 0x2}, {0x4cf, 0x52f, 0x2}, {0x530, 0x557, 0x27}, {0x558, 0x109f, 0x1}, {0x10c6, 0x10c8, 0x2}, {0x10c9, 0x10cc, 0x1}, {0x10ce, 0x139f, 0x1}, {0x13f6, 0x1dff, 0x1}, {0x1e01, 0x1e95, 0x2}, {0x1e96, 0x1e9d, 0x1}, {0x1e9f, 0x1eff, 0x2}, {0x1f00, 0x1f07, 0x1}, {0x1f10, 0x1f17, 0x1}, {0x1f1e, 0x1f27, 0x1}, {0x1f30, 0x1f37, 0x1}, {0x1f40, 0x1f47, 0x1}, {0x1f4e, 0x1f58, 0x1}, {0x1f5a, 0x1f60, 0x2}, {0x1f61, 0x1f67, 0x1}, {0x1f70, 0x1fb7, 0x1}, {0x1fbc, 0x1fc7, 0x1}, {0x1fcc, 0x1fd7, 0x1}, {0x1fdc, 0x1fe7, 0x1}, {0x1fed, 0x1ff7, 0x1}, {0x1ffc, 0x2101, 0x1}, {0x2103, 0x2106, 0x1}, {0x2108, 0x210a, 0x1}, {0x210e, 0x210f, 0x1}, {0x2113, 0x2114, 0x1}, {0x2116, 0x2118, 0x1}, {0x211e, 0x2123, 0x1}, {0x2125, 0x2129, 0x2}, {0x212e, 0x212f, 0x1}, {0x2134, 0x213d, 0x1}, {0x2140, 0x2144, 0x1}, {0x2146, 0x215f, 0x1}, {0x2170, 0x2182, 0x1}, {0x2184, 0x24b5, 0x1}, {0x24d0, 0x2bff, 0x1}, {0x2c2f, 0x2c5f, 0x1}, {0x2c61, 0x2c65, 0x4}, {0x2c66, 0x2c6c, 0x2}, {0x2c71, 0x2c73, 0x2}, {0x2c74, 0x2c76, 0x2}, {0x2c77, 0x2c7d, 0x1}, {0x2c81, 0x2ce3, 0x2}, {0x2ce4, 0x2cea, 0x1}, {0x2cec, 0x2cee, 0x2}, {0x2cef, 0x2cf1, 0x1}, {0x2cf3, 0xa63f, 0x1}, {0xa641, 0xa66d, 0x2}, {0xa66e, 0xa67f, 0x1}, {0xa681, 0xa69b, 0x2}, {0xa69c, 0xa721, 0x1}, {0xa723, 0xa72f, 0x2}, {0xa730, 0xa731, 0x1}, {0xa733, 0xa76f, 0x2}, {0xa770, 0xa778, 0x1}, {0xa77a, 0xa77c, 0x2}, {0xa77f, 0xa787, 0x2}, {0xa788, 0xa78a, 0x1}, {0xa78c, 0xa78e, 0x2}, {0xa78f, 0xa793, 0x2}, {0xa794, 0xa795, 0x1}, {0xa797, 0xa7a9, 0x2}, {0xa7af, 0xa7b5, 0x6}, {0xa7b7, 0xff20, 0x1}, {0xff3b, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x103ff, 0x1}, {0x10428, 0x104af, 0x1}, {0x104d4, 0x10c7f, 0x1}, {0x10cb3, 0x1189f, 0x1}, {0x118c0, 0x1d3ff, 0x1}, {0x1d41a, 0x1d433, 0x1}, {0x1d44e, 0x1d467, 0x1}, {0x1d482, 0x1d49b, 0x1}, {0x1d49d, 0x1d4a0, 0x3}, {0x1d4a1, 0x1d4a3, 0x2}, {0x1d4a4, 0x1d4a7, 0x3}, {0x1d4a8, 0x1d4ad, 0x5}, {0x1d4b6, 0x1d4cf, 0x1}, {0x1d4ea, 0x1d503, 0x1}, {0x1d506, 0x1d50b, 0x5}, {0x1d50c, 0x1d515, 0x9}, {0x1d51d, 0x1d537, 0x1}, {0x1d53a, 0x1d53f, 0x5}, {0x1d545, 0x1d547, 0x2}, {0x1d548, 0x1d549, 0x1}, {0x1d551, 0x1d56b, 0x1}, {0x1d586, 0x1d59f, 0x1}, {0x1d5ba, 0x1d5d3, 0x1}, {0x1d5ee, 0x1d607, 0x1}, {0x1d622, 0x1d63b, 0x1}, {0x1d656, 0x1d66f, 0x1}, {0x1d68a, 0x1d6a7, 0x1}, {0x1d6c1, 0x1d6e1, 0x1}, {0x1d6fb, 0x1d71b, 0x1}, {0x1d735, 0x1d755, 0x1}, {0x1d76f, 0x1d78f, 0x1}, {0x1d7a9, 0x1d7c9, 0x1}, {0x1d7cb, 0x1e8ff, 0x1}, {0x1e922, 0x1f12f, 0x1}, {0x1f14a, 0x1f14f, 0x1}, {0x1f16a, 0x1f16f, 0x1}, {0x1f18a, 0x10ffff, 0x1}}, 4}
	uppercaseYes = &unicode.RangeTable{[]unicode.Range16{{0x41, 0x5a, 0x1}, {0xc0, 0xd6, 0x1}, {0xd8, 0xde, 0x1}, {0x100, 0x136, 0x2}, {0x139, 0x147, 0x2}, {0x14a, 0x178, 0x2}, {0x179, 0x17d, 0x2}, {0x181, 0x182, 0x1}, {0x184, 0x186, 0x2}, {0x187, 0x189, 0x2}, {0x18a, 0x18b, 0x1}, {0x18e, 0x191, 0x1}, {0x193, 0x194, 0x1}, {0x196, 0x198, 0x1}, {0x19c, 0x19d, 0x1}, {0x19f, 0x1a0, 0x1}, {0x1a2, 0x1a6, 0x2}, {0x1a7, 0x1a9, 0x2}, {0x1ac, 0x1ae, 0x2}, {0x1af, 0x1b1, 0x2}, {0x1b2, 0x1b3, 0x1}, {0x1b5, 0x1b7, 0x2}, {0x1b8, 0x1bc, 0x4}, {0x1c4, 0x1cd, 0x3}, {0x1cf, 0x1db, 0x2}, {0x1de, 0x1ee, 0x2}, {0x1f1, 0x1f4, 0x3}, {0x1f6, 0x1f8, 0x1}, {0x1fa, 0x232, 0x2}, {0x23a, 0x23b, 0x1}, {0x23d, 0x23e, 0x1}, {0x241, 0x243, 0x2}, {0x244, 0x246, 0x1}, {0x248, 0x24e, 0x2}, {0x370, 0x372, 0x2}, {0x376, 0x37f, 0x9}, {0x386, 0x388, 0x2}, {0x389, 0x38a, 0x1}, {0x38c, 0x38e, 0x2}, {0x38f, 0x391, 0x2}, {0x392, 0x3a1, 0x1}, {0x3a3, 0x3ab, 0x1}, {0x3cf, 0x3d2, 0x3}, {0x3d3, 0x3d4, 0x1}, {0x3d8, 0x3ee, 0x2}, {0x3f4, 0x3f7, 0x3}, {0x3f9, 0x3fa, 0x1}, {0x3fd, 0x42f, 0x1}, {0x460, 0x480, 0x2}, {0x48a, 0x4c0, 0x2}, {0x4c1, 0x4cd, 0x2}, {0x4d0, 0x52e, 0x2}, {0x531, 0x556, 0x1}, {0x10a0, 0x10c5, 0x1}, {0x10c7, 0x10cd, 0x6}, {0x13a0, 0x13f5, 0x1}, {0x1e00, 0x1e94, 0x2}, {0x1e9e, 0x1efe, 0x2}, {0x1f08, 0x1f0f, 0x1}, {0x1f18, 0x1f1d, 0x1}, {0x1f28, 0x1f2f, 0x1}, {0x1f38, 0x1f3f, 0x1}, {0x1f48, 0x1f4d, 0x1}, {0x1f59, 0x1f5f, 0x2}, {0x1f68, 0x1f6f, 0x1}, {0x1fb8, 0x1fbb, 0x1}, {0x1fc8, 0x1fcb, 0x1}, {0x1fd8, 0x1fdb, 0x1}, {0x1fe8, 0x1fec, 0x1}, {0x1ff8, 0x1ffb, 0x1}, {0x2102, 0x2107, 0x5}, {0x210b, 0x210d, 0x1}, {0x2110, 0x2112, 0x1}, {0x2115, 0x2119, 0x4}, {0x211a, 0x211d, 0x1}, {0x2124, 0x212a, 0x2}, {0x212b, 0x212d, 0x1}, {0x2130, 0x2133, 0x1}, {0x213e, 0x213f, 0x1}, {0x2145, 0x2160, 0x1b}, {0x2161, 0x216f, 0x1}, {0x2183, 0x24b6, 0x333}, {0x24b7, 0x24cf, 0x1}, {0x2c00, 0x2c2e, 0x1}, {0x2c60, 0x2c62, 0x2}, {0x2c63, 0x2c64, 0x1}, {0x2c67, 0x2c6d, 0x2}, {0x2c6e, 0x2c70, 0x1}, {0x2c72, 0x2c75, 0x3}, {0x2c7e, 0x2c80, 0x1}, {0x2c82, 0x2ce2, 0x2}, {0x2ceb, 0x2ced, 0x2}, {0x2cf2, 0xa640, 0x794e}, {0xa642, 0xa66c, 0x2}, {0xa680, 0xa69a, 0x2}, {0xa722, 0xa72e, 0x2}, {0xa732, 0xa76e, 0x2}, {0xa779, 0xa77d, 0x2}, {0xa77e, 0xa786, 0x2}, {0xa78b, 0xa78d, 0x2}, {0xa790, 0xa792, 0x2}, {0xa796, 0xa7aa, 0x2}, {0xa7ab, 0xa7ae, 0x1}, {0xa7b0, 0xa7b4, 0x1}, {0xa7b6, 0xff21, 0x576b}, {0xff22, 0xff3a, 0x1}}, []unicode.Range32{{0x10400, 0x10427, 0x1}, {0x104b0, 0x104d3, 0x1}, {0x10c80, 0x10cb2, 0x1}, {0x118a0, 0x118bf, 0x1}, {0x1d400, 0x1d419, 0x1}, {0x1d434, 0x1d44d, 0x1}, {0x1d468, 0x1d481, 0x1}, {0x1d49c, 0x1d49e, 0x2}, {0x1d49f, 0x1d4a5, 0x3}, {0x1d4a6, 0x1d4a9, 0x3}, {0x1d4aa, 0x1d4ac, 0x1}, {0x1d4ae, 0x1d4b5, 0x1}, {0x1d4d0, 0x1d4e9, 0x1}, {0x1d504, 0x1d505, 0x1}, {0x1d507, 0x1d50a, 0x1}, {0x1d50d, 0x1d514, 0x1}, {0x1d516, 0x1d51c, 0x1}, {0x1d538, 0x1d539, 0x1}, {0x1d53b, 0x1d53e, 0x1}, {0x1d540, 0x1d544, 0x1}, {0x1d546, 0x1d54a, 0x4}, {0x1d54b, 0x1d550, 0x1}, {0x1d56c, 0x1d585, 0x1}, {0x1d5a0, 0x1d5b9, 0x1}, {0x1d5d4, 0x1d5ed, 0x1}, {0x1d608, 0x1d621, 0x1}, {0x1d63c, 0x1d655, 0x1}, {0x1d670, 0x1d689, 0x1}, {0x1d6a8, 0x1d6c0, 0x1}, {0x1d6e2, 0x1d6fa, 0x1}, {0x1d71c, 0x1d734, 0x1}, {0x1d756, 0x1d76e, 0x1}, {0x1d790, 0x1d7a8, 0x1}, {0x1d7ca, 0x1e900, 0x1136}, {0x1e901, 0x1e921, 0x1}, {0x1f130, 0x1f149, 0x1}, {0x1f150, 0x1f169, 0x1}, {0x1f170, 0x1f189, 0x1}}, 3}
)