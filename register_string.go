// Code generated by "stringer -type=Register -linecomment"; DO NOT EDIT.

package zydis

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RegisterNone-0]
	_ = x[RegisterAL-1]
	_ = x[RegisterCL-2]
	_ = x[RegisterDL-3]
	_ = x[RegisterBL-4]
	_ = x[RegisterAH-5]
	_ = x[RegisterCH-6]
	_ = x[RegisterDH-7]
	_ = x[RegisterBH-8]
	_ = x[RegisterSPL-9]
	_ = x[RegisterBPL-10]
	_ = x[RegisterSIL-11]
	_ = x[RegisterDIL-12]
	_ = x[RegisterR8B-13]
	_ = x[RegisterR9B-14]
	_ = x[RegisterR10B-15]
	_ = x[RegisterR11B-16]
	_ = x[RegisterR12B-17]
	_ = x[RegisterR13B-18]
	_ = x[RegisterR14B-19]
	_ = x[RegisterR15B-20]
	_ = x[RegisterAX-21]
	_ = x[RegisterCX-22]
	_ = x[RegisterDX-23]
	_ = x[RegisterBX-24]
	_ = x[RegisterSP-25]
	_ = x[RegisterBP-26]
	_ = x[RegisterSI-27]
	_ = x[RegisterDI-28]
	_ = x[RegisterR8W-29]
	_ = x[RegisterR9W-30]
	_ = x[RegisterR10W-31]
	_ = x[RegisterR11W-32]
	_ = x[RegisterR12W-33]
	_ = x[RegisterR13W-34]
	_ = x[RegisterR14W-35]
	_ = x[RegisterR15W-36]
	_ = x[RegisterEAX-37]
	_ = x[RegisterECX-38]
	_ = x[RegisterEDX-39]
	_ = x[RegisterEBX-40]
	_ = x[RegisterESP-41]
	_ = x[RegisterEBP-42]
	_ = x[RegisterESI-43]
	_ = x[RegisterEDI-44]
	_ = x[RegisterR8D-45]
	_ = x[RegisterR9D-46]
	_ = x[RegisterR10D-47]
	_ = x[RegisterR11D-48]
	_ = x[RegisterR12D-49]
	_ = x[RegisterR13D-50]
	_ = x[RegisterR14D-51]
	_ = x[RegisterR15D-52]
	_ = x[RegisterRAX-53]
	_ = x[RegisterRCX-54]
	_ = x[RegisterRDX-55]
	_ = x[RegisterRBX-56]
	_ = x[RegisterRSP-57]
	_ = x[RegisterRBP-58]
	_ = x[RegisterRSI-59]
	_ = x[RegisterRDI-60]
	_ = x[RegisterR8-61]
	_ = x[RegisterR9-62]
	_ = x[RegisterR10-63]
	_ = x[RegisterR11-64]
	_ = x[RegisterR12-65]
	_ = x[RegisterR13-66]
	_ = x[RegisterR14-67]
	_ = x[RegisterR15-68]
	_ = x[RegisterST0-69]
	_ = x[RegisterST1-70]
	_ = x[RegisterST2-71]
	_ = x[RegisterST3-72]
	_ = x[RegisterST4-73]
	_ = x[RegisterST5-74]
	_ = x[RegisterST6-75]
	_ = x[RegisterST7-76]
	_ = x[RegisterX87Control-77]
	_ = x[RegisterX87Status-78]
	_ = x[RegisterX87Tag-79]
	_ = x[RegisterMM0-80]
	_ = x[RegisterMM1-81]
	_ = x[RegisterMM2-82]
	_ = x[RegisterMM3-83]
	_ = x[RegisterMM4-84]
	_ = x[RegisterMM5-85]
	_ = x[RegisterMM6-86]
	_ = x[RegisterMM7-87]
	_ = x[RegisterXMM0-88]
	_ = x[RegisterXMM1-89]
	_ = x[RegisterXMM2-90]
	_ = x[RegisterXMM3-91]
	_ = x[RegisterXMM4-92]
	_ = x[RegisterXMM5-93]
	_ = x[RegisterXMM6-94]
	_ = x[RegisterXMM7-95]
	_ = x[RegisterXMM8-96]
	_ = x[RegisterXMM9-97]
	_ = x[RegisterXMM10-98]
	_ = x[RegisterXMM11-99]
	_ = x[RegisterXMM12-100]
	_ = x[RegisterXMM13-101]
	_ = x[RegisterXMM14-102]
	_ = x[RegisterXMM15-103]
	_ = x[RegisterXMM16-104]
	_ = x[RegisterXMM17-105]
	_ = x[RegisterXMM18-106]
	_ = x[RegisterXMM19-107]
	_ = x[RegisterXMM20-108]
	_ = x[RegisterXMM21-109]
	_ = x[RegisterXMM22-110]
	_ = x[RegisterXMM23-111]
	_ = x[RegisterXMM24-112]
	_ = x[RegisterXMM25-113]
	_ = x[RegisterXMM26-114]
	_ = x[RegisterXMM27-115]
	_ = x[RegisterXMM28-116]
	_ = x[RegisterXMM29-117]
	_ = x[RegisterXMM30-118]
	_ = x[RegisterXMM31-119]
	_ = x[RegisterYMM0-120]
	_ = x[RegisterYMM1-121]
	_ = x[RegisterYMM2-122]
	_ = x[RegisterYMM3-123]
	_ = x[RegisterYMM4-124]
	_ = x[RegisterYMM5-125]
	_ = x[RegisterYMM6-126]
	_ = x[RegisterYMM7-127]
	_ = x[RegisterYMM8-128]
	_ = x[RegisterYMM9-129]
	_ = x[RegisterYMM10-130]
	_ = x[RegisterYMM11-131]
	_ = x[RegisterYMM12-132]
	_ = x[RegisterYMM13-133]
	_ = x[RegisterYMM14-134]
	_ = x[RegisterYMM15-135]
	_ = x[RegisterYMM16-136]
	_ = x[RegisterYMM17-137]
	_ = x[RegisterYMM18-138]
	_ = x[RegisterYMM19-139]
	_ = x[RegisterYMM20-140]
	_ = x[RegisterYMM21-141]
	_ = x[RegisterYMM22-142]
	_ = x[RegisterYMM23-143]
	_ = x[RegisterYMM24-144]
	_ = x[RegisterYMM25-145]
	_ = x[RegisterYMM26-146]
	_ = x[RegisterYMM27-147]
	_ = x[RegisterYMM28-148]
	_ = x[RegisterYMM29-149]
	_ = x[RegisterYMM30-150]
	_ = x[RegisterYMM31-151]
	_ = x[RegisterZMM0-152]
	_ = x[RegisterZMM1-153]
	_ = x[RegisterZMM2-154]
	_ = x[RegisterZMM3-155]
	_ = x[RegisterZMM4-156]
	_ = x[RegisterZMM5-157]
	_ = x[RegisterZMM6-158]
	_ = x[RegisterZMM7-159]
	_ = x[RegisterZMM8-160]
	_ = x[RegisterZMM9-161]
	_ = x[RegisterZMM10-162]
	_ = x[RegisterZMM11-163]
	_ = x[RegisterZMM12-164]
	_ = x[RegisterZMM13-165]
	_ = x[RegisterZMM14-166]
	_ = x[RegisterZMM15-167]
	_ = x[RegisterZMM16-168]
	_ = x[RegisterZMM17-169]
	_ = x[RegisterZMM18-170]
	_ = x[RegisterZMM19-171]
	_ = x[RegisterZMM20-172]
	_ = x[RegisterZMM21-173]
	_ = x[RegisterZMM22-174]
	_ = x[RegisterZMM23-175]
	_ = x[RegisterZMM24-176]
	_ = x[RegisterZMM25-177]
	_ = x[RegisterZMM26-178]
	_ = x[RegisterZMM27-179]
	_ = x[RegisterZMM28-180]
	_ = x[RegisterZMM29-181]
	_ = x[RegisterZMM30-182]
	_ = x[RegisterZMM31-183]
	_ = x[RegisterFlags-184]
	_ = x[RegisterEFlags-185]
	_ = x[RegisterRFlags-186]
	_ = x[RegisterIP-187]
	_ = x[RegisterEIP-188]
	_ = x[RegisterRIP-189]
	_ = x[RegisterES-190]
	_ = x[RegisterCS-191]
	_ = x[RegisterSS-192]
	_ = x[RegisterDS-193]
	_ = x[RegisterFS-194]
	_ = x[RegisterGS-195]
	_ = x[RegisterGDTR-196]
	_ = x[RegisterLDTR-197]
	_ = x[RegisterIDTR-198]
	_ = x[RegisterTR-199]
	_ = x[RegisterTR0-200]
	_ = x[RegisterTR1-201]
	_ = x[RegisterTR2-202]
	_ = x[RegisterTR3-203]
	_ = x[RegisterTR4-204]
	_ = x[RegisterTR5-205]
	_ = x[RegisterTR6-206]
	_ = x[RegisterTR7-207]
	_ = x[RegisterCR0-208]
	_ = x[RegisterCR1-209]
	_ = x[RegisterCR2-210]
	_ = x[RegisterCR3-211]
	_ = x[RegisterCR4-212]
	_ = x[RegisterCR5-213]
	_ = x[RegisterCR6-214]
	_ = x[RegisterCR7-215]
	_ = x[RegisterCR8-216]
	_ = x[RegisterCR9-217]
	_ = x[RegisterCR10-218]
	_ = x[RegisterCR11-219]
	_ = x[RegisterCR12-220]
	_ = x[RegisterCR13-221]
	_ = x[RegisterCR14-222]
	_ = x[RegisterCR15-223]
	_ = x[RegisterDR0-224]
	_ = x[RegisterDR1-225]
	_ = x[RegisterDR2-226]
	_ = x[RegisterDR3-227]
	_ = x[RegisterDR4-228]
	_ = x[RegisterDR5-229]
	_ = x[RegisterDR6-230]
	_ = x[RegisterDR7-231]
	_ = x[RegisterDR8-232]
	_ = x[RegisterDR9-233]
	_ = x[RegisterDR10-234]
	_ = x[RegisterDR11-235]
	_ = x[RegisterDR12-236]
	_ = x[RegisterDR13-237]
	_ = x[RegisterDR14-238]
	_ = x[RegisterDR15-239]
	_ = x[RegisterK0-240]
	_ = x[RegisterK1-241]
	_ = x[RegisterK2-242]
	_ = x[RegisterK3-243]
	_ = x[RegisterK4-244]
	_ = x[RegisterK5-245]
	_ = x[RegisterK6-246]
	_ = x[RegisterK7-247]
	_ = x[RegisterBND0-248]
	_ = x[RegisterBND1-249]
	_ = x[RegisterBND2-250]
	_ = x[RegisterBND3-251]
	_ = x[RegisterBNDCfg-252]
	_ = x[RegisterBNDStatus-253]
	_ = x[RegisterMXCSR-254]
	_ = x[RegisterPKRU-255]
	_ = x[RegisterXCR0-256]
}

const _Register_name = "NoneALCLDLBLAHCHDHBHSPLBPLSILDILR8BR9BR10BR11BR12BR13BR14BR15BAXCXDXBXSPBPSIDIR8WR9WR10WR11WR12WR13WR14WR15WEAXECXEDXEBXESPEBPESIEDIR8DR9DR10DR11DR12DR13DR14DR15DRAXRCXRDXRBXRSPRBPRSIRDIR8R9R10R11R12R13R14R15ST0ST1ST2ST3ST4ST5ST6ST7X87ControlX87StatusX87TagMM0MM1MM2MM3MM4MM5MM6MM7XMM0XMM1XMM2XMM3XMM4XMM5XMM6XMM7XMM8XMM9XMM10XMM11XMM12XMM13XMM14XMM15XMM16XMM17XMM18XMM19XMM20XMM21XMM22XMM23XMM24XMM25XMM26XMM27XMM28XMM29XMM30XMM31YMM0YMM1YMM2YMM3YMM4YMM5YMM6YMM7YMM8YMM9YMM10YMM11YMM12YMM13YMM14YMM15YMM16YMM17YMM18YMM19YMM20YMM21YMM22YMM23YMM24YMM25YMM26YMM27YMM28YMM29YMM30YMM31ZMM0ZMM1ZMM2ZMM3ZMM4ZMM5ZMM6ZMM7ZMM8ZMM9ZMM10ZMM11ZMM12ZMM13ZMM14ZMM15ZMM16ZMM17ZMM18ZMM19ZMM20ZMM21ZMM22ZMM23ZMM24ZMM25ZMM26ZMM27ZMM28ZMM29ZMM30ZMM31FlagsEFlagsRFlagsIPEIPRIPESCSSSDSFSGSGDTRLDTRIDTRTRTR0TR1TR2TR3TR4TR5TR6TR7CR0CR1CR2CR3CR4CR5CR6CR7CR8CR9CR10CR11CR12CR13CR14CR15DR0DR1DR2DR3DR4DR5DR6DR7DR8DR9DR10DR11DR12DR13DR14DR15K0K1K2K3K4K5K6K7BND0BND1BND2BND3BNDCfgBNDStatusMXCSRPKRUXCR0"

var _Register_index = [...]uint16{0, 4, 6, 8, 10, 12, 14, 16, 18, 20, 23, 26, 29, 32, 35, 38, 42, 46, 50, 54, 58, 62, 64, 66, 68, 70, 72, 74, 76, 78, 81, 84, 88, 92, 96, 100, 104, 108, 111, 114, 117, 120, 123, 126, 129, 132, 135, 138, 142, 146, 150, 154, 158, 162, 165, 168, 171, 174, 177, 180, 183, 186, 188, 190, 193, 196, 199, 202, 205, 208, 211, 214, 217, 220, 223, 226, 229, 232, 242, 251, 257, 260, 263, 266, 269, 272, 275, 278, 281, 285, 289, 293, 297, 301, 305, 309, 313, 317, 321, 326, 331, 336, 341, 346, 351, 356, 361, 366, 371, 376, 381, 386, 391, 396, 401, 406, 411, 416, 421, 426, 431, 435, 439, 443, 447, 451, 455, 459, 463, 467, 471, 476, 481, 486, 491, 496, 501, 506, 511, 516, 521, 526, 531, 536, 541, 546, 551, 556, 561, 566, 571, 576, 581, 585, 589, 593, 597, 601, 605, 609, 613, 617, 621, 626, 631, 636, 641, 646, 651, 656, 661, 666, 671, 676, 681, 686, 691, 696, 701, 706, 711, 716, 721, 726, 731, 736, 742, 748, 750, 753, 756, 758, 760, 762, 764, 766, 768, 772, 776, 780, 782, 785, 788, 791, 794, 797, 800, 803, 806, 809, 812, 815, 818, 821, 824, 827, 830, 833, 836, 840, 844, 848, 852, 856, 860, 863, 866, 869, 872, 875, 878, 881, 884, 887, 890, 894, 898, 902, 906, 910, 914, 916, 918, 920, 922, 924, 926, 928, 930, 934, 938, 942, 946, 952, 961, 966, 970, 974}

func (i Register) String() string {
	if i < 0 || i >= Register(len(_Register_index)-1) {
		return "Register(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Register_name[_Register_index[i]:_Register_index[i+1]]
}
