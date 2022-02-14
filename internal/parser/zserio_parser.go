// Code generated from /private/var/tmp/_bazel_ignas.anikevicius/c3129ead79ca509099c649d74caf832e/sandbox/darwin-sandbox/2595/execroot/__main__/internal/parser/ZserioParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ZserioParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 113, 760, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55, 
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9, 
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65, 
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4, 
	71, 9, 71, 3, 2, 5, 2, 144, 10, 2, 3, 2, 7, 2, 147, 10, 2, 12, 2, 14, 2, 
	150, 11, 2, 3, 2, 7, 2, 153, 10, 2, 12, 2, 14, 2, 156, 11, 2, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 164, 10, 3, 12, 3, 14, 3, 167, 11, 3, 
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 177, 10, 4, 12, 4, 
	14, 4, 180, 11, 4, 3, 4, 3, 4, 5, 4, 184, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 
	5, 5, 190, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 
	3, 6, 3, 6, 5, 6, 203, 10, 6, 3, 7, 3, 7, 5, 7, 207, 10, 7, 3, 8, 3, 8, 
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 220, 10, 9, 
	12, 9, 14, 9, 223, 11, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 240, 10, 
	12, 3, 12, 5, 12, 243, 10, 12, 3, 12, 3, 12, 7, 12, 247, 10, 12, 12, 12, 
	14, 12, 250, 11, 12, 3, 12, 7, 12, 253, 10, 12, 12, 12, 14, 12, 256, 11, 
	12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13, 262, 10, 13, 3, 13, 5, 13, 265, 
	10, 13, 3, 13, 5, 13, 268, 10, 13, 3, 13, 3, 13, 5, 13, 272, 10, 13, 3, 
	13, 5, 13, 275, 10, 13, 3, 13, 5, 13, 278, 10, 13, 3, 13, 3, 13, 3, 14, 
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 5, 16, 292, 
	10, 16, 3, 16, 5, 16, 295, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 300, 10, 
	16, 3, 17, 3, 17, 5, 17, 304, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 5, 21, 320, 
	10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 327, 10, 21, 12, 21, 
	14, 21, 330, 11, 21, 3, 21, 5, 21, 333, 10, 21, 3, 21, 7, 21, 336, 10, 
	21, 12, 21, 14, 21, 339, 11, 21, 3, 21, 3, 21, 3, 21, 3, 22, 6, 22, 345, 
	10, 22, 13, 22, 14, 22, 346, 3, 22, 5, 22, 350, 10, 22, 3, 22, 3, 22, 3, 
	23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 361, 10, 24, 3, 24, 
	3, 24, 3, 25, 3, 25, 5, 25, 367, 10, 25, 3, 26, 3, 26, 3, 26, 5, 26, 372, 
	10, 26, 3, 26, 5, 26, 375, 10, 26, 3, 26, 3, 26, 7, 26, 379, 10, 26, 12, 
	26, 14, 26, 382, 11, 26, 3, 26, 7, 26, 385, 10, 26, 12, 26, 14, 26, 388, 
	11, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 
	3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 403, 10, 28, 12, 28, 14, 28, 406, 11, 
	28, 3, 28, 5, 28, 409, 10, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 
	5, 29, 417, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 
	30, 426, 10, 30, 12, 30, 14, 30, 429, 11, 30, 3, 30, 5, 30, 432, 10, 30, 
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 440, 10, 31, 3, 32, 3, 
	32, 3, 32, 5, 32, 445, 10, 32, 3, 32, 3, 32, 5, 32, 449, 10, 32, 3, 32, 
	3, 32, 7, 32, 453, 10, 32, 12, 32, 14, 32, 456, 11, 32, 3, 32, 5, 32, 459, 
	10, 32, 3, 32, 5, 32, 462, 10, 32, 3, 32, 3, 32, 3, 32, 3, 33, 5, 33, 468, 
	10, 33, 3, 33, 3, 33, 3, 33, 5, 33, 473, 10, 33, 3, 33, 3, 33, 3, 34, 3, 
	34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 
	3, 37, 6, 37, 490, 10, 37, 13, 37, 14, 37, 491, 3, 37, 3, 37, 3, 37, 3, 
	38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 7, 39, 505, 10, 39, 
	12, 39, 14, 39, 508, 11, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 
	3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 524, 10, 
	41, 12, 41, 14, 41, 527, 11, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 
	42, 3, 42, 3, 42, 3, 43, 5, 43, 538, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43, 
	3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 
	46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 
	3, 48, 7, 48, 566, 10, 48, 12, 48, 14, 48, 569, 11, 48, 3, 48, 3, 48, 3, 
	49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 7, 50, 580, 10, 50, 12, 50, 
	14, 50, 583, 11, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 7, 51, 591, 
	10, 51, 12, 51, 14, 51, 594, 11, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 
	3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 630, 10, 
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 
	5, 54, 642, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 
	7, 54, 683, 10, 54, 12, 54, 14, 54, 686, 11, 54, 3, 55, 3, 55, 3, 56, 3, 
	56, 3, 57, 3, 57, 3, 57, 5, 57, 695, 10, 57, 5, 57, 697, 10, 57, 3, 58, 
	3, 58, 3, 58, 5, 58, 702, 10, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 
	59, 3, 59, 3, 59, 5, 59, 712, 10, 59, 3, 60, 3, 60, 3, 60, 7, 60, 717, 
	10, 60, 12, 60, 14, 60, 720, 11, 60, 3, 61, 3, 61, 3, 61, 3, 61, 7, 61, 
	726, 10, 61, 12, 61, 14, 61, 729, 11, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 
	62, 5, 62, 736, 10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 
	3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 
	69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 2, 3, 106, 72, 2, 4, 6, 8, 10, 12, 
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 
	86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 
	118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 2, 13, 4, 2, 
	66, 66, 78, 78, 6, 2, 5, 5, 21, 21, 26, 26, 32, 32, 4, 2, 8, 8, 22, 23, 
	4, 2, 21, 21, 26, 26, 5, 2, 11, 12, 15, 15, 20, 20, 4, 2, 10, 10, 24, 24, 
	3, 2, 102, 109, 4, 2, 55, 58, 81, 84, 3, 2, 88, 96, 4, 2, 35, 35, 54, 54, 
	3, 2, 45, 47, 2, 790, 2, 143, 3, 2, 2, 2, 4, 159, 3, 2, 2, 2, 6, 170, 3, 
	2, 2, 2, 8, 189, 3, 2, 2, 2, 10, 202, 3, 2, 2, 2, 12, 206, 3, 2, 2, 2, 
	14, 208, 3, 2, 2, 2, 16, 215, 3, 2, 2, 2, 18, 227, 3, 2, 2, 2, 20, 231, 
	3, 2, 2, 2, 22, 236, 3, 2, 2, 2, 24, 261, 3, 2, 2, 2, 26, 281, 3, 2, 2, 
	2, 28, 287, 3, 2, 2, 2, 30, 291, 3, 2, 2, 2, 32, 301, 3, 2, 2, 2, 34, 307, 
	3, 2, 2, 2, 36, 310, 3, 2, 2, 2, 38, 313, 3, 2, 2, 2, 40, 316, 3, 2, 2, 
	2, 42, 344, 3, 2, 2, 2, 44, 353, 3, 2, 2, 2, 46, 357, 3, 2, 2, 2, 48, 364, 
	3, 2, 2, 2, 50, 368, 3, 2, 2, 2, 52, 392, 3, 2, 2, 2, 54, 395, 3, 2, 2, 
	2, 56, 413, 3, 2, 2, 2, 58, 418, 3, 2, 2, 2, 60, 436, 3, 2, 2, 2, 62, 441, 
	3, 2, 2, 2, 64, 467, 3, 2, 2, 2, 66, 476, 3, 2, 2, 2, 68, 479, 3, 2, 2, 
	2, 70, 482, 3, 2, 2, 2, 72, 485, 3, 2, 2, 2, 74, 496, 3, 2, 2, 2, 76, 500, 
	3, 2, 2, 2, 78, 512, 3, 2, 2, 2, 80, 519, 3, 2, 2, 2, 82, 531, 3, 2, 2, 
	2, 84, 537, 3, 2, 2, 2, 86, 544, 3, 2, 2, 2, 88, 551, 3, 2, 2, 2, 90, 553, 
	3, 2, 2, 2, 92, 555, 3, 2, 2, 2, 94, 561, 3, 2, 2, 2, 96, 572, 3, 2, 2, 
	2, 98, 575, 3, 2, 2, 2, 100, 586, 3, 2, 2, 2, 102, 597, 3, 2, 2, 2, 104, 
	599, 3, 2, 2, 2, 106, 629, 3, 2, 2, 2, 108, 687, 3, 2, 2, 2, 110, 689, 
	3, 2, 2, 2, 112, 696, 3, 2, 2, 2, 114, 698, 3, 2, 2, 2, 116, 711, 3, 2, 
	2, 2, 118, 713, 3, 2, 2, 2, 120, 721, 3, 2, 2, 2, 122, 735, 3, 2, 2, 2, 
	124, 737, 3, 2, 2, 2, 126, 741, 3, 2, 2, 2, 128, 743, 3, 2, 2, 2, 130, 
	745, 3, 2, 2, 2, 132, 749, 3, 2, 2, 2, 134, 751, 3, 2, 2, 2, 136, 753, 
	3, 2, 2, 2, 138, 755, 3, 2, 2, 2, 140, 757, 3, 2, 2, 2, 142, 144, 5, 4, 
	3, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 148, 3, 2, 2, 2, 
	145, 147, 5, 6, 4, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 
	146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 154, 3, 2, 2, 2, 150, 148, 
	3, 2, 2, 2, 151, 153, 5, 8, 5, 2, 152, 151, 3, 2, 2, 2, 153, 156, 3, 2, 
	2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 
	156, 154, 3, 2, 2, 2, 157, 158, 7, 2, 2, 3, 158, 3, 3, 2, 2, 2, 159, 160, 
	7, 63, 2, 2, 160, 165, 5, 110, 56, 2, 161, 162, 7, 9, 2, 2, 162, 164, 5, 
	110, 56, 2, 163, 161, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 
	2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 
	168, 169, 7, 31, 2, 2, 169, 5, 3, 2, 2, 2, 170, 171, 7, 51, 2, 2, 171, 
	172, 5, 110, 56, 2, 172, 178, 7, 9, 2, 2, 173, 174, 5, 110, 56, 2, 174, 
	175, 7, 9, 2, 2, 175, 177, 3, 2, 2, 2, 176, 173, 3, 2, 2, 2, 177, 180, 
	3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 183, 3, 2, 
	2, 2, 180, 178, 3, 2, 2, 2, 181, 184, 5, 110, 56, 2, 182, 184, 7, 23, 2, 
	2, 183, 181, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 
	186, 7, 31, 2, 2, 186, 7, 3, 2, 2, 2, 187, 190, 5, 12, 7, 2, 188, 190, 
	5, 10, 6, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 9, 3, 2, 
	2, 2, 191, 203, 5, 20, 11, 2, 192, 203, 5, 22, 12, 2, 193, 203, 5, 40, 
	21, 2, 194, 203, 5, 50, 26, 2, 195, 203, 5, 54, 28, 2, 196, 203, 5, 58, 
	30, 2, 197, 203, 5, 62, 32, 2, 198, 203, 5, 72, 37, 2, 199, 203, 5, 76, 
	39, 2, 200, 203, 5, 80, 41, 2, 201, 203, 5, 104, 53, 2, 202, 191, 3, 2, 
	2, 2, 202, 192, 3, 2, 2, 2, 202, 193, 3, 2, 2, 2, 202, 194, 3, 2, 2, 2, 
	202, 195, 3, 2, 2, 2, 202, 196, 3, 2, 2, 2, 202, 197, 3, 2, 2, 2, 202, 
	198, 3, 2, 2, 2, 202, 199, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 201, 
	3, 2, 2, 2, 203, 11, 3, 2, 2, 2, 204, 207, 5, 14, 8, 2, 205, 207, 5, 16, 
	9, 2, 206, 204, 3, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 13, 3, 2, 2, 2, 
	208, 209, 7, 40, 2, 2, 209, 210, 5, 114, 58, 2, 210, 211, 5, 110, 56, 2, 
	211, 212, 7, 4, 2, 2, 212, 213, 5, 106, 54, 2, 213, 214, 7, 31, 2, 2, 214, 
	15, 3, 2, 2, 2, 215, 216, 7, 69, 2, 2, 216, 217, 5, 110, 56, 2, 217, 221, 
	7, 13, 2, 2, 218, 220, 5, 18, 10, 2, 219, 218, 3, 2, 2, 2, 220, 223, 3, 
	2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 3, 2, 2, 
	2, 223, 221, 3, 2, 2, 2, 224, 225, 7, 28, 2, 2, 225, 226, 7, 31, 2, 2, 
	226, 17, 3, 2, 2, 2, 227, 228, 7, 68, 2, 2, 228, 229, 5, 106, 54, 2, 229, 
	230, 7, 31, 2, 2, 230, 19, 3, 2, 2, 2, 231, 232, 7, 79, 2, 2, 232, 233, 
	5, 112, 57, 2, 233, 234, 5, 110, 56, 2, 234, 235, 7, 31, 2, 2, 235, 21, 
	3, 2, 2, 2, 236, 237, 7, 77, 2, 2, 237, 239, 5, 110, 56, 2, 238, 240, 5, 
	98, 50, 2, 239, 238, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 242, 3, 2, 
	2, 2, 241, 243, 5, 94, 48, 2, 242, 241, 3, 2, 2, 2, 242, 243, 3, 2, 2, 
	2, 243, 244, 3, 2, 2, 2, 244, 248, 7, 13, 2, 2, 245, 247, 5, 24, 13, 2, 
	246, 245, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 
	249, 3, 2, 2, 2, 249, 254, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251, 253, 
	5, 86, 44, 2, 252, 251, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 
	2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 257, 3, 2, 2, 2, 256, 254, 3, 2, 2, 
	2, 257, 258, 7, 28, 2, 2, 258, 259, 7, 31, 2, 2, 259, 23, 3, 2, 2, 2, 260, 
	262, 5, 26, 14, 2, 261, 260, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 264, 
	3, 2, 2, 2, 263, 265, 5, 28, 15, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 
	2, 2, 2, 265, 267, 3, 2, 2, 2, 266, 268, 7, 62, 2, 2, 267, 266, 3, 2, 2, 
	2, 267, 268, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 271, 5, 30, 16, 2, 
	270, 272, 5, 34, 18, 2, 271, 270, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 
	274, 3, 2, 2, 2, 273, 275, 5, 36, 19, 2, 274, 273, 3, 2, 2, 2, 274, 275, 
	3, 2, 2, 2, 275, 277, 3, 2, 2, 2, 276, 278, 5, 38, 20, 2, 277, 276, 3, 
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 7, 31, 2, 
	2, 280, 25, 3, 2, 2, 2, 281, 282, 7, 34, 2, 2, 282, 283, 7, 18, 2, 2, 283, 
	284, 5, 106, 54, 2, 284, 285, 7, 30, 2, 2, 285, 286, 7, 6, 2, 2, 286, 27, 
	3, 2, 2, 2, 287, 288, 5, 106, 54, 2, 288, 289, 7, 6, 2, 2, 289, 29, 3, 
	2, 2, 2, 290, 292, 7, 64, 2, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 
	2, 292, 294, 3, 2, 2, 2, 293, 295, 7, 50, 2, 2, 294, 293, 3, 2, 2, 2, 294, 
	295, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 5, 114, 58, 2, 297, 299, 
	5, 110, 56, 2, 298, 300, 5, 32, 17, 2, 299, 298, 3, 2, 2, 2, 299, 300, 
	3, 2, 2, 2, 300, 31, 3, 2, 2, 2, 301, 303, 7, 14, 2, 2, 302, 304, 5, 106, 
	54, 2, 303, 302, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 
	305, 306, 7, 29, 2, 2, 306, 33, 3, 2, 2, 2, 307, 308, 7, 4, 2, 2, 308, 
	309, 5, 106, 54, 2, 309, 35, 3, 2, 2, 2, 310, 311, 7, 49, 2, 2, 311, 312, 
	5, 106, 54, 2, 312, 37, 3, 2, 2, 2, 313, 314, 7, 6, 2, 2, 314, 315, 5, 
	106, 54, 2, 315, 39, 3, 2, 2, 2, 316, 317, 7, 39, 2, 2, 317, 319, 5, 110, 
	56, 2, 318, 320, 5, 98, 50, 2, 319, 318, 3, 2, 2, 2, 319, 320, 3, 2, 2, 
	2, 320, 321, 3, 2, 2, 2, 321, 322, 5, 94, 48, 2, 322, 323, 7, 61, 2, 2, 
	323, 324, 5, 106, 54, 2, 324, 328, 7, 13, 2, 2, 325, 327, 5, 42, 22, 2, 
	326, 325, 3, 2, 2, 2, 327, 330, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 
	329, 3, 2, 2, 2, 329, 332, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 331, 333, 
	5, 46, 24, 2, 332, 331, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 337, 3, 
	2, 2, 2, 334, 336, 5, 86, 44, 2, 335, 334, 3, 2, 2, 2, 336, 339, 3, 2, 
	2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 340, 3, 2, 2, 2, 
	339, 337, 3, 2, 2, 2, 340, 341, 7, 28, 2, 2, 341, 342, 7, 31, 2, 2, 342, 
	41, 3, 2, 2, 2, 343, 345, 5, 44, 23, 2, 344, 343, 3, 2, 2, 2, 345, 346, 
	3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 349, 3, 2, 
	2, 2, 348, 350, 5, 48, 25, 2, 349, 348, 3, 2, 2, 2, 349, 350, 3, 2, 2, 
	2, 350, 351, 3, 2, 2, 2, 351, 352, 7, 31, 2, 2, 352, 43, 3, 2, 2, 2, 353, 
	354, 7, 38, 2, 2, 354, 355, 5, 106, 54, 2, 355, 356, 7, 6, 2, 2, 356, 45, 
	3, 2, 2, 2, 357, 358, 7, 41, 2, 2, 358, 360, 7, 6, 2, 2, 359, 361, 5, 48, 
	25, 2, 360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 
	362, 363, 7, 31, 2, 2, 363, 47, 3, 2, 2, 2, 364, 366, 5, 30, 16, 2, 365, 
	367, 5, 38, 20, 2, 366, 365, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 49, 
	3, 2, 2, 2, 368, 369, 7, 85, 2, 2, 369, 371, 5, 110, 56, 2, 370, 372, 5, 
	98, 50, 2, 371, 370, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 374, 3, 2, 
	2, 2, 373, 375, 5, 94, 48, 2, 374, 373, 3, 2, 2, 2, 374, 375, 3, 2, 2, 
	2, 375, 376, 3, 2, 2, 2, 376, 380, 7, 13, 2, 2, 377, 379, 5, 52, 27, 2, 
	378, 377, 3, 2, 2, 2, 379, 382, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 
	381, 3, 2, 2, 2, 381, 386, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 383, 385, 
	5, 86, 44, 2, 384, 383, 3, 2, 2, 2, 385, 388, 3, 2, 2, 2, 386, 384, 3, 
	2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 389, 3, 2, 2, 2, 388, 386, 3, 2, 2, 
	2, 389, 390, 7, 28, 2, 2, 390, 391, 7, 31, 2, 2, 391, 51, 3, 2, 2, 2, 392, 
	393, 5, 48, 25, 2, 393, 394, 7, 31, 2, 2, 394, 53, 3, 2, 2, 2, 395, 396, 
	7, 42, 2, 2, 396, 397, 5, 114, 58, 2, 397, 398, 5, 110, 56, 2, 398, 399, 
	7, 13, 2, 2, 399, 404, 5, 56, 29, 2, 400, 401, 7, 7, 2, 2, 401, 403, 5, 
	56, 29, 2, 402, 400, 3, 2, 2, 2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 
	2, 2, 404, 405, 3, 2, 2, 2, 405, 408, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 
	407, 409, 7, 7, 2, 2, 408, 407, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 
	410, 3, 2, 2, 2, 410, 411, 7, 28, 2, 2, 411, 412, 7, 31, 2, 2, 412, 55, 
	3, 2, 2, 2, 413, 416, 5, 110, 56, 2, 414, 415, 7, 4, 2, 2, 415, 417, 5, 
	106, 54, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 57, 3, 2, 
	2, 2, 418, 419, 7, 37, 2, 2, 419, 420, 5, 114, 58, 2, 420, 421, 5, 110, 
	56, 2, 421, 422, 7, 13, 2, 2, 422, 427, 5, 60, 31, 2, 423, 424, 7, 7, 2, 
	2, 424, 426, 5, 60, 31, 2, 425, 423, 3, 2, 2, 2, 426, 429, 3, 2, 2, 2, 
	427, 425, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 431, 3, 2, 2, 2, 429, 
	427, 3, 2, 2, 2, 430, 432, 7, 7, 2, 2, 431, 430, 3, 2, 2, 2, 431, 432, 
	3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 7, 28, 2, 2, 434, 435, 7, 31, 
	2, 2, 435, 59, 3, 2, 2, 2, 436, 439, 5, 110, 56, 2, 437, 438, 7, 4, 2, 
	2, 438, 440, 5, 106, 54, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 
	440, 61, 3, 2, 2, 2, 441, 442, 7, 73, 2, 2, 442, 444, 5, 110, 56, 2, 443, 
	445, 5, 98, 50, 2, 444, 443, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 448, 
	3, 2, 2, 2, 446, 447, 7, 86, 2, 2, 447, 449, 5, 110, 56, 2, 448, 446, 3, 
	2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 454, 7, 13, 2, 
	2, 451, 453, 5, 64, 33, 2, 452, 451, 3, 2, 2, 2, 453, 456, 3, 2, 2, 2, 
	454, 452, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 458, 3, 2, 2, 2, 456, 
	454, 3, 2, 2, 2, 457, 459, 5, 66, 34, 2, 458, 457, 3, 2, 2, 2, 458, 459, 
	3, 2, 2, 2, 459, 461, 3, 2, 2, 2, 460, 462, 5, 70, 36, 2, 461, 460, 3, 
	2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 464, 7, 28, 2, 
	2, 464, 465, 7, 31, 2, 2, 465, 63, 3, 2, 2, 2, 466, 468, 7, 74, 2, 2, 467, 
	466, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 
	5, 114, 58, 2, 470, 472, 5, 110, 56, 2, 471, 473, 5, 68, 35, 2, 472, 471, 
	3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 475, 7, 31, 
	2, 2, 475, 65, 3, 2, 2, 2, 476, 477, 5, 68, 35, 2, 477, 478, 7, 31, 2, 
	2, 478, 67, 3, 2, 2, 2, 479, 480, 7, 71, 2, 2, 480, 481, 5, 106, 54, 2, 
	481, 69, 3, 2, 2, 2, 482, 483, 7, 75, 2, 2, 483, 484, 7, 31, 2, 2, 484, 
	71, 3, 2, 2, 2, 485, 486, 7, 72, 2, 2, 486, 487, 5, 110, 56, 2, 487, 489, 
	7, 13, 2, 2, 488, 490, 5, 74, 38, 2, 489, 488, 3, 2, 2, 2, 490, 491, 3, 
	2, 2, 2, 491, 489, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 493, 3, 2, 2, 
	2, 493, 494, 7, 28, 2, 2, 494, 495, 7, 31, 2, 2, 495, 73, 3, 2, 2, 2, 496, 
	497, 5, 114, 58, 2, 497, 498, 5, 110, 56, 2, 498, 499, 7, 31, 2, 2, 499, 
	75, 3, 2, 2, 2, 500, 501, 7, 70, 2, 2, 501, 502, 5, 110, 56, 2, 502, 506, 
	7, 13, 2, 2, 503, 505, 5, 78, 40, 2, 504, 503, 3, 2, 2, 2, 505, 508, 3, 
	2, 2, 2, 506, 504, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 509, 3, 2, 2, 
	2, 508, 506, 3, 2, 2, 2, 509, 510, 7, 28, 2, 2, 510, 511, 7, 31, 2, 2, 
	511, 77, 3, 2, 2, 2, 512, 513, 5, 112, 57, 2, 513, 514, 5, 110, 56, 2, 
	514, 515, 7, 18, 2, 2, 515, 516, 5, 112, 57, 2, 516, 517, 7, 30, 2, 2, 
	517, 518, 7, 31, 2, 2, 518, 79, 3, 2, 2, 2, 519, 520, 7, 65, 2, 2, 520, 
	521, 5, 110, 56, 2, 521, 525, 7, 13, 2, 2, 522, 524, 5, 82, 42, 2, 523, 
	522, 3, 2, 2, 2, 524, 527, 3, 2, 2, 2, 525, 523, 3, 2, 2, 2, 525, 526, 
	3, 2, 2, 2, 526, 528, 3, 2, 2, 2, 527, 525, 3, 2, 2, 2, 528, 529, 7, 28, 
	2, 2, 529, 530, 7, 31, 2, 2, 530, 81, 3, 2, 2, 2, 531, 532, 5, 84, 43, 
	2, 532, 533, 5, 112, 57, 2, 533, 534, 5, 110, 56, 2, 534, 535, 7, 31, 2, 
	2, 535, 83, 3, 2, 2, 2, 536, 538, 9, 2, 2, 2, 537, 536, 3, 2, 2, 2, 537, 
	538, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 540, 7, 80, 2, 2, 540, 541, 
	7, 18, 2, 2, 541, 542, 5, 106, 54, 2, 542, 543, 7, 30, 2, 2, 543, 85, 3, 
	2, 2, 2, 544, 545, 7, 48, 2, 2, 545, 546, 5, 88, 45, 2, 546, 547, 5, 90, 
	46, 2, 547, 548, 7, 18, 2, 2, 548, 549, 7, 30, 2, 2, 549, 550, 5, 92, 47, 
	2, 550, 87, 3, 2, 2, 2, 551, 552, 5, 112, 57, 2, 552, 89, 3, 2, 2, 2, 553, 
	554, 5, 110, 56, 2, 554, 91, 3, 2, 2, 2, 555, 556, 7, 13, 2, 2, 556, 557, 
	7, 67, 2, 2, 557, 558, 5, 106, 54, 2, 558, 559, 7, 31, 2, 2, 559, 560, 
	7, 28, 2, 2, 560, 93, 3, 2, 2, 2, 561, 562, 7, 18, 2, 2, 562, 567, 5, 96, 
	49, 2, 563, 564, 7, 7, 2, 2, 564, 566, 5, 96, 49, 2, 565, 563, 3, 2, 2, 
	2, 566, 569, 3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568, 
	570, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 570, 571, 7, 30, 2, 2, 571, 95, 
	3, 2, 2, 2, 572, 573, 5, 112, 57, 2, 573, 574, 5, 110, 56, 2, 574, 97, 
	3, 2, 2, 2, 575, 576, 7, 20, 2, 2, 576, 581, 5, 110, 56, 2, 577, 578, 7, 
	7, 2, 2, 578, 580, 5, 110, 56, 2, 579, 577, 3, 2, 2, 2, 580, 583, 3, 2, 
	2, 2, 581, 579, 3, 2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 584, 3, 2, 2, 2, 
	583, 581, 3, 2, 2, 2, 584, 585, 7, 12, 2, 2, 585, 99, 3, 2, 2, 2, 586, 
	587, 7, 20, 2, 2, 587, 592, 5, 102, 52, 2, 588, 589, 7, 7, 2, 2, 589, 591, 
	5, 102, 52, 2, 590, 588, 3, 2, 2, 2, 591, 594, 3, 2, 2, 2, 592, 590, 3, 
	2, 2, 2, 592, 593, 3, 2, 2, 2, 593, 595, 3, 2, 2, 2, 594, 592, 3, 2, 2, 
	2, 595, 596, 7, 12, 2, 2, 596, 101, 3, 2, 2, 2, 597, 598, 5, 112, 57, 2, 
	598, 103, 3, 2, 2, 2, 599, 600, 7, 53, 2, 2, 600, 601, 5, 112, 57, 2, 601, 
	602, 5, 110, 56, 2, 602, 603, 7, 31, 2, 2, 603, 105, 3, 2, 2, 2, 604, 605, 
	8, 54, 1, 2, 605, 606, 7, 18, 2, 2, 606, 607, 5, 106, 54, 2, 607, 608, 
	7, 30, 2, 2, 608, 630, 3, 2, 2, 2, 609, 610, 7, 59, 2, 2, 610, 611, 7, 
	18, 2, 2, 611, 612, 5, 106, 54, 2, 612, 613, 7, 30, 2, 2, 613, 630, 3, 
	2, 2, 2, 614, 615, 7, 87, 2, 2, 615, 616, 7, 18, 2, 2, 616, 617, 5, 106, 
	54, 2, 617, 618, 7, 30, 2, 2, 618, 630, 3, 2, 2, 2, 619, 620, 7, 60, 2, 
	2, 620, 621, 7, 18, 2, 2, 621, 622, 5, 106, 54, 2, 622, 623, 7, 30, 2, 
	2, 623, 630, 3, 2, 2, 2, 624, 625, 9, 3, 2, 2, 625, 630, 5, 106, 54, 17, 
	626, 630, 5, 108, 55, 2, 627, 630, 7, 52, 2, 2, 628, 630, 5, 110, 56, 2, 
	629, 604, 3, 2, 2, 2, 629, 609, 3, 2, 2, 2, 629, 614, 3, 2, 2, 2, 629, 
	619, 3, 2, 2, 2, 629, 624, 3, 2, 2, 2, 629, 626, 3, 2, 2, 2, 629, 627, 
	3, 2, 2, 2, 629, 628, 3, 2, 2, 2, 630, 684, 3, 2, 2, 2, 631, 632, 12, 16, 
	2, 2, 632, 633, 9, 4, 2, 2, 633, 683, 5, 106, 54, 17, 634, 635, 12, 15, 
	2, 2, 635, 636, 9, 5, 2, 2, 636, 683, 5, 106, 54, 16, 637, 641, 12, 14, 
	2, 2, 638, 642, 7, 19, 2, 2, 639, 640, 7, 12, 2, 2, 640, 642, 7, 12, 2, 
	2, 641, 638, 3, 2, 2, 2, 641, 639, 3, 2, 2, 2, 642, 643, 3, 2, 2, 2, 643, 
	683, 5, 106, 54, 15, 644, 645, 12, 13, 2, 2, 645, 646, 9, 6, 2, 2, 646, 
	683, 5, 106, 54, 14, 647, 648, 12, 12, 2, 2, 648, 649, 9, 7, 2, 2, 649, 
	683, 5, 106, 54, 13, 650, 651, 12, 11, 2, 2, 651, 652, 7, 3, 2, 2, 652, 
	683, 5, 106, 54, 12, 653, 654, 12, 10, 2, 2, 654, 655, 7, 33, 2, 2, 655, 
	683, 5, 106, 54, 11, 656, 657, 12, 9, 2, 2, 657, 658, 7, 25, 2, 2, 658, 
	683, 5, 106, 54, 10, 659, 660, 12, 8, 2, 2, 660, 661, 7, 16, 2, 2, 661, 
	683, 5, 106, 54, 9, 662, 663, 12, 7, 2, 2, 663, 664, 7, 17, 2, 2, 664, 
	683, 5, 106, 54, 8, 665, 666, 12, 6, 2, 2, 666, 667, 7, 27, 2, 2, 667, 
	668, 5, 106, 54, 2, 668, 669, 7, 6, 2, 2, 669, 670, 5, 106, 54, 6, 670, 
	683, 3, 2, 2, 2, 671, 672, 12, 23, 2, 2, 672, 673, 7, 18, 2, 2, 673, 683, 
	7, 30, 2, 2, 674, 675, 12, 22, 2, 2, 675, 676, 7, 14, 2, 2, 676, 677, 5, 
	106, 54, 2, 677, 678, 7, 29, 2, 2, 678, 683, 3, 2, 2, 2, 679, 680, 12, 
	21, 2, 2, 680, 681, 7, 9, 2, 2, 681, 683, 5, 110, 56, 2, 682, 631, 3, 2, 
	2, 2, 682, 634, 3, 2, 2, 2, 682, 637, 3, 2, 2, 2, 682, 644, 3, 2, 2, 2, 
	682, 647, 3, 2, 2, 2, 682, 650, 3, 2, 2, 2, 682, 653, 3, 2, 2, 2, 682, 
	656, 3, 2, 2, 2, 682, 659, 3, 2, 2, 2, 682, 662, 3, 2, 2, 2, 682, 665, 
	3, 2, 2, 2, 682, 671, 3, 2, 2, 2, 682, 674, 3, 2, 2, 2, 682, 679, 3, 2, 
	2, 2, 683, 686, 3, 2, 2, 2, 684, 682, 3, 2, 2, 2, 684, 685, 3, 2, 2, 2, 
	685, 107, 3, 2, 2, 2, 686, 684, 3, 2, 2, 2, 687, 688, 9, 8, 2, 2, 688, 
	109, 3, 2, 2, 2, 689, 690, 7, 110, 2, 2, 690, 111, 3, 2, 2, 2, 691, 697, 
	5, 116, 59, 2, 692, 694, 5, 118, 60, 2, 693, 695, 5, 100, 51, 2, 694, 693, 
	3, 2, 2, 2, 694, 695, 3, 2, 2, 2, 695, 697, 3, 2, 2, 2, 696, 691, 3, 2, 
	2, 2, 696, 692, 3, 2, 2, 2, 697, 113, 3, 2, 2, 2, 698, 701, 5, 112, 57, 
	2, 699, 702, 5, 120, 61, 2, 700, 702, 5, 124, 63, 2, 701, 699, 3, 2, 2, 
	2, 701, 700, 3, 2, 2, 2, 701, 702, 3, 2, 2, 2, 702, 115, 3, 2, 2, 2, 703, 
	712, 5, 126, 64, 2, 704, 712, 5, 128, 65, 2, 705, 712, 5, 130, 66, 2, 706, 
	712, 5, 132, 67, 2, 707, 712, 5, 134, 68, 2, 708, 712, 5, 136, 69, 2, 709, 
	712, 5, 138, 70, 2, 710, 712, 5, 140, 71, 2, 711, 703, 3, 2, 2, 2, 711, 
	704, 3, 2, 2, 2, 711, 705, 3, 2, 2, 2, 711, 706, 3, 2, 2, 2, 711, 707, 
	3, 2, 2, 2, 711, 708, 3, 2, 2, 2, 711, 709, 3, 2, 2, 2, 711, 710, 3, 2, 
	2, 2, 712, 117, 3, 2, 2, 2, 713, 718, 5, 110, 56, 2, 714, 715, 7, 9, 2, 
	2, 715, 717, 5, 110, 56, 2, 716, 714, 3, 2, 2, 2, 717, 720, 3, 2, 2, 2, 
	718, 716, 3, 2, 2, 2, 718, 719, 3, 2, 2, 2, 719, 119, 3, 2, 2, 2, 720, 
	718, 3, 2, 2, 2, 721, 722, 7, 18, 2, 2, 722, 727, 5, 122, 62, 2, 723, 724, 
	7, 7, 2, 2, 724, 726, 5, 122, 62, 2, 725, 723, 3, 2, 2, 2, 726, 729, 3, 
	2, 2, 2, 727, 725, 3, 2, 2, 2, 727, 728, 3, 2, 2, 2, 728, 730, 3, 2, 2, 
	2, 729, 727, 3, 2, 2, 2, 730, 731, 7, 30, 2, 2, 731, 121, 3, 2, 2, 2, 732, 
	733, 7, 43, 2, 2, 733, 736, 5, 110, 56, 2, 734, 736, 5, 106, 54, 2, 735, 
	732, 3, 2, 2, 2, 735, 734, 3, 2, 2, 2, 736, 123, 3, 2, 2, 2, 737, 738, 
	7, 20, 2, 2, 738, 739, 5, 106, 54, 2, 739, 740, 7, 12, 2, 2, 740, 125, 
	3, 2, 2, 2, 741, 742, 9, 9, 2, 2, 742, 127, 3, 2, 2, 2, 743, 744, 9, 10, 
	2, 2, 744, 129, 3, 2, 2, 2, 745, 746, 9, 11, 2, 2, 746, 747, 7, 6, 2, 2, 
	747, 748, 7, 109, 2, 2, 748, 131, 3, 2, 2, 2, 749, 750, 9, 11, 2, 2, 750, 
	133, 3, 2, 2, 2, 751, 752, 7, 36, 2, 2, 752, 135, 3, 2, 2, 2, 753, 754, 
	7, 76, 2, 2, 754, 137, 3, 2, 2, 2, 755, 756, 9, 12, 2, 2, 756, 139, 3, 
	2, 2, 2, 757, 758, 7, 44, 2, 2, 758, 141, 3, 2, 2, 2, 69, 143, 148, 154, 
	165, 178, 183, 189, 202, 206, 221, 239, 242, 248, 254, 261, 264, 267, 271, 
	274, 277, 291, 294, 299, 303, 319, 328, 332, 337, 346, 349, 360, 366, 371, 
	374, 380, 386, 404, 408, 416, 427, 431, 439, 444, 448, 454, 458, 461, 467, 
	472, 491, 506, 525, 537, 567, 581, 592, 629, 641, 682, 684, 694, 696, 701, 
	711, 718, 727, 735,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'&'", "'='", "'!'", "':'", "','", "'/'", "'.'", "'=='", "'>='", "'>'", 
	"'{'", "'['", "'<='", "'&&'", "'||'", "'('", "'<<'", "'<'", "'-'", "'%'", 
	"'*'", "'!='", "'|'", "'+'", "'?'", "'}'", "']'", "')'", "';'", "'~'", 
	"'^'", "'align'", "'bit'", "'bool'", "'bitmask'", "'case'", "'choice'", 
	"'const'", "'default'", "'enum'", "'explicit'", "'extern'", "'float16'", 
	"'float32'", "'float64'", "'function'", "'if'", "'implicit'", "'import'", 
	"'@index'", "'instantiate'", "'int'", "'int16'", "'int32'", "'int64'", 
	"'int8'", "'lengthof'", "'numbits'", "'on'", "'optional'", "'package'", 
	"'packed'", "'pubsub'", "'publish'", "'return'", "'rule'", "'rule_group'", 
	"'service'", "'sql'", "'sql_database'", "'sql_table'", "'sql_virtual'", 
	"'sql_without_rowid'", "'string'", "'struct'", "'subscribe'", "'subtype'", 
	"'topic'", "'uint16'", "'uint32'", "'uint64'", "'uint8'", "'union'", "'using'", 
	"'valueof'", "'varint'", "'varint16'", "'varint32'", "'varint64'", "'varsize'", 
	"'varuint'", "'varuint16'", "'varuint32'", "'varuint64'",
}
var symbolicNames = []string{
	"", "AND", "ASSIGN", "BANG", "COLON", "COMMA", "DIVIDE", "DOT", "EQ", "GE", 
	"GT", "LBRACE", "LBRACKET", "LE", "LOGICAL_AND", "LOGICAL_OR", "LPAREN", 
	"LSHIFT", "LT", "MINUS", "MODULO", "MULTIPLY", "NE", "OR", "PLUS", "QUESTIONMARK", 
	"RBRACE", "RBRACKET", "RPAREN", "SEMICOLON", "TILDE", "XOR", "ALIGN", "BIT_FIELD", 
	"BOOL", "BITMASK", "CASE", "CHOICE", "CONST", "DEFAULT", "ENUM", "EXPLICIT", 
	"EXTERN", "FLOAT16", "FLOAT32", "FLOAT64", "FUNCTION", "IF", "IMPLICIT", 
	"IMPORT", "INDEX", "INSTANTIATE", "INT_FIELD", "INT16", "INT32", "INT64", 
	"INT8", "LENGTHOF", "NUMBITS", "ON", "OPTIONAL", "PACKAGE", "PACKED", "PUBSUB", 
	"PUBLISH", "RETURN", "RULE", "RULE_GROUP", "SERVICE", "SQL", "SQL_DATABASE", 
	"SQL_TABLE", "SQL_VIRTUAL", "SQL_WITHOUT_ROWID", "STRING", "STRUCTURE", 
	"SUBSCRIBE", "SUBTYPE", "TOPIC", "UINT16", "UINT32", "UINT64", "UINT8", 
	"UNION", "USING", "VALUEOF", "VARINT", "VARINT16", "VARINT32", "VARINT64", 
	"VARSIZE", "VARUINT", "VARUINT16", "VARUINT32", "VARUINT64", "WS", "DOC_COMMENT", 
	"MARKDOWN_COMMENT", "BLOCK_COMMENT", "LINE_COMMENT", "BOOL_LITERAL", "STRING_LITERAL", 
	"BINARY_LITERAL", "OCTAL_LITERAL", "HEXADECIMAL_LITERAL", "DOUBLE_LITERAL", 
	"FLOAT_LITERAL", "DECIMAL_LITERAL", "ID", "INVALID_STRING_LITERAL", "INVALID_TOKEN", 
	"RSHIFT",
}

var ruleNames = []string{
	"packageDeclaration", "packageNameDefinition", "importDeclaration", "languageDirective", 
	"typeDeclaration", "symbolDefinition", "constDefinition", "ruleGroupDefinition", 
	"ruleDefinition", "subtypeDeclaration", "structureDeclaration", "structureFieldDefinition", 
	"fieldAlignment", "fieldOffset", "fieldTypeId", "fieldArrayRange", "fieldInitializer", 
	"fieldOptionalClause", "fieldConstraint", "choiceDeclaration", "choiceCases", 
	"choiceCase", "choiceDefault", "choiceFieldDefinition", "unionDeclaration", 
	"unionFieldDefinition", "enumDeclaration", "enumItem", "bitmaskDeclaration", 
	"bitmaskValue", "sqlTableDeclaration", "sqlTableFieldDefinition", "sqlConstraintDefinition", 
	"sqlConstraint", "sqlWithoutRowId", "sqlDatabaseDefinition", "sqlDatabaseFieldDefinition", 
	"serviceDefinition", "serviceMethodDefinition", "pubsubDefinition", "pubsubMessageDefinition", 
	"topicDefinition", "functionDefinition", "functionType", "functionName", 
	"functionBody", "typeParameters", "parameterDefinition", "templateParameters", 
	"templateArguments", "templateArgument", "instantiateDeclaration", "expression", 
	"literal", "id", "typeReference", "typeInstantiation", "builtinType", "qualifiedName", 
	"typeArguments", "typeArgument", "dynamicLengthArgument", "intType", "varintType", 
	"fixedBitFieldType", "dynamicBitFieldType", "boolType", "stringType", "floatType", 
	"externType",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ZserioParser struct {
	*antlr.BaseParser
}

func NewZserioParser(input antlr.TokenStream) *ZserioParser {
	this := new(ZserioParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ZserioParser.g4"

	return this
}

// ZserioParser tokens.
const (
	ZserioParserEOF = antlr.TokenEOF
	ZserioParserAND = 1
	ZserioParserASSIGN = 2
	ZserioParserBANG = 3
	ZserioParserCOLON = 4
	ZserioParserCOMMA = 5
	ZserioParserDIVIDE = 6
	ZserioParserDOT = 7
	ZserioParserEQ = 8
	ZserioParserGE = 9
	ZserioParserGT = 10
	ZserioParserLBRACE = 11
	ZserioParserLBRACKET = 12
	ZserioParserLE = 13
	ZserioParserLOGICAL_AND = 14
	ZserioParserLOGICAL_OR = 15
	ZserioParserLPAREN = 16
	ZserioParserLSHIFT = 17
	ZserioParserLT = 18
	ZserioParserMINUS = 19
	ZserioParserMODULO = 20
	ZserioParserMULTIPLY = 21
	ZserioParserNE = 22
	ZserioParserOR = 23
	ZserioParserPLUS = 24
	ZserioParserQUESTIONMARK = 25
	ZserioParserRBRACE = 26
	ZserioParserRBRACKET = 27
	ZserioParserRPAREN = 28
	ZserioParserSEMICOLON = 29
	ZserioParserTILDE = 30
	ZserioParserXOR = 31
	ZserioParserALIGN = 32
	ZserioParserBIT_FIELD = 33
	ZserioParserBOOL = 34
	ZserioParserBITMASK = 35
	ZserioParserCASE = 36
	ZserioParserCHOICE = 37
	ZserioParserCONST = 38
	ZserioParserDEFAULT = 39
	ZserioParserENUM = 40
	ZserioParserEXPLICIT = 41
	ZserioParserEXTERN = 42
	ZserioParserFLOAT16 = 43
	ZserioParserFLOAT32 = 44
	ZserioParserFLOAT64 = 45
	ZserioParserFUNCTION = 46
	ZserioParserIF = 47
	ZserioParserIMPLICIT = 48
	ZserioParserIMPORT = 49
	ZserioParserINDEX = 50
	ZserioParserINSTANTIATE = 51
	ZserioParserINT_FIELD = 52
	ZserioParserINT16 = 53
	ZserioParserINT32 = 54
	ZserioParserINT64 = 55
	ZserioParserINT8 = 56
	ZserioParserLENGTHOF = 57
	ZserioParserNUMBITS = 58
	ZserioParserON = 59
	ZserioParserOPTIONAL = 60
	ZserioParserPACKAGE = 61
	ZserioParserPACKED = 62
	ZserioParserPUBSUB = 63
	ZserioParserPUBLISH = 64
	ZserioParserRETURN = 65
	ZserioParserRULE = 66
	ZserioParserRULE_GROUP = 67
	ZserioParserSERVICE = 68
	ZserioParserSQL = 69
	ZserioParserSQL_DATABASE = 70
	ZserioParserSQL_TABLE = 71
	ZserioParserSQL_VIRTUAL = 72
	ZserioParserSQL_WITHOUT_ROWID = 73
	ZserioParserSTRING = 74
	ZserioParserSTRUCTURE = 75
	ZserioParserSUBSCRIBE = 76
	ZserioParserSUBTYPE = 77
	ZserioParserTOPIC = 78
	ZserioParserUINT16 = 79
	ZserioParserUINT32 = 80
	ZserioParserUINT64 = 81
	ZserioParserUINT8 = 82
	ZserioParserUNION = 83
	ZserioParserUSING = 84
	ZserioParserVALUEOF = 85
	ZserioParserVARINT = 86
	ZserioParserVARINT16 = 87
	ZserioParserVARINT32 = 88
	ZserioParserVARINT64 = 89
	ZserioParserVARSIZE = 90
	ZserioParserVARUINT = 91
	ZserioParserVARUINT16 = 92
	ZserioParserVARUINT32 = 93
	ZserioParserVARUINT64 = 94
	ZserioParserWS = 95
	ZserioParserDOC_COMMENT = 96
	ZserioParserMARKDOWN_COMMENT = 97
	ZserioParserBLOCK_COMMENT = 98
	ZserioParserLINE_COMMENT = 99
	ZserioParserBOOL_LITERAL = 100
	ZserioParserSTRING_LITERAL = 101
	ZserioParserBINARY_LITERAL = 102
	ZserioParserOCTAL_LITERAL = 103
	ZserioParserHEXADECIMAL_LITERAL = 104
	ZserioParserDOUBLE_LITERAL = 105
	ZserioParserFLOAT_LITERAL = 106
	ZserioParserDECIMAL_LITERAL = 107
	ZserioParserID = 108
	ZserioParserINVALID_STRING_LITERAL = 109
	ZserioParserINVALID_TOKEN = 110
	ZserioParserRSHIFT = 111
)

// ZserioParser rules.
const (
	ZserioParserRULE_packageDeclaration = 0
	ZserioParserRULE_packageNameDefinition = 1
	ZserioParserRULE_importDeclaration = 2
	ZserioParserRULE_languageDirective = 3
	ZserioParserRULE_typeDeclaration = 4
	ZserioParserRULE_symbolDefinition = 5
	ZserioParserRULE_constDefinition = 6
	ZserioParserRULE_ruleGroupDefinition = 7
	ZserioParserRULE_ruleDefinition = 8
	ZserioParserRULE_subtypeDeclaration = 9
	ZserioParserRULE_structureDeclaration = 10
	ZserioParserRULE_structureFieldDefinition = 11
	ZserioParserRULE_fieldAlignment = 12
	ZserioParserRULE_fieldOffset = 13
	ZserioParserRULE_fieldTypeId = 14
	ZserioParserRULE_fieldArrayRange = 15
	ZserioParserRULE_fieldInitializer = 16
	ZserioParserRULE_fieldOptionalClause = 17
	ZserioParserRULE_fieldConstraint = 18
	ZserioParserRULE_choiceDeclaration = 19
	ZserioParserRULE_choiceCases = 20
	ZserioParserRULE_choiceCase = 21
	ZserioParserRULE_choiceDefault = 22
	ZserioParserRULE_choiceFieldDefinition = 23
	ZserioParserRULE_unionDeclaration = 24
	ZserioParserRULE_unionFieldDefinition = 25
	ZserioParserRULE_enumDeclaration = 26
	ZserioParserRULE_enumItem = 27
	ZserioParserRULE_bitmaskDeclaration = 28
	ZserioParserRULE_bitmaskValue = 29
	ZserioParserRULE_sqlTableDeclaration = 30
	ZserioParserRULE_sqlTableFieldDefinition = 31
	ZserioParserRULE_sqlConstraintDefinition = 32
	ZserioParserRULE_sqlConstraint = 33
	ZserioParserRULE_sqlWithoutRowId = 34
	ZserioParserRULE_sqlDatabaseDefinition = 35
	ZserioParserRULE_sqlDatabaseFieldDefinition = 36
	ZserioParserRULE_serviceDefinition = 37
	ZserioParserRULE_serviceMethodDefinition = 38
	ZserioParserRULE_pubsubDefinition = 39
	ZserioParserRULE_pubsubMessageDefinition = 40
	ZserioParserRULE_topicDefinition = 41
	ZserioParserRULE_functionDefinition = 42
	ZserioParserRULE_functionType = 43
	ZserioParserRULE_functionName = 44
	ZserioParserRULE_functionBody = 45
	ZserioParserRULE_typeParameters = 46
	ZserioParserRULE_parameterDefinition = 47
	ZserioParserRULE_templateParameters = 48
	ZserioParserRULE_templateArguments = 49
	ZserioParserRULE_templateArgument = 50
	ZserioParserRULE_instantiateDeclaration = 51
	ZserioParserRULE_expression = 52
	ZserioParserRULE_literal = 53
	ZserioParserRULE_id = 54
	ZserioParserRULE_typeReference = 55
	ZserioParserRULE_typeInstantiation = 56
	ZserioParserRULE_builtinType = 57
	ZserioParserRULE_qualifiedName = 58
	ZserioParserRULE_typeArguments = 59
	ZserioParserRULE_typeArgument = 60
	ZserioParserRULE_dynamicLengthArgument = 61
	ZserioParserRULE_intType = 62
	ZserioParserRULE_varintType = 63
	ZserioParserRULE_fixedBitFieldType = 64
	ZserioParserRULE_dynamicBitFieldType = 65
	ZserioParserRULE_boolType = 66
	ZserioParserRULE_stringType = 67
	ZserioParserRULE_floatType = 68
	ZserioParserRULE_externType = 69
)

// IPackageDeclarationContext is an interface to support dynamic dispatch.
type IPackageDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageDeclarationContext differentiates from other interfaces.
	IsPackageDeclarationContext()
}

type PackageDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDeclarationContext() *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_packageDeclaration
	return p
}

func (*PackageDeclarationContext) IsPackageDeclarationContext() {}

func NewPackageDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_packageDeclaration

	return p
}

func (s *PackageDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDeclarationContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZserioParserEOF, 0)
}

func (s *PackageDeclarationContext) PackageNameDefinition() IPackageNameDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageNameDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageNameDefinitionContext)
}

func (s *PackageDeclarationContext) AllImportDeclaration() []IImportDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem())
	var tst = make([]IImportDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclarationContext)
		}
	}

	return tst
}

func (s *PackageDeclarationContext) ImportDeclaration(i int) IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *PackageDeclarationContext) AllLanguageDirective() []ILanguageDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILanguageDirectiveContext)(nil)).Elem())
	var tst = make([]ILanguageDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILanguageDirectiveContext)
		}
	}

	return tst
}

func (s *PackageDeclarationContext) LanguageDirective(i int) ILanguageDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILanguageDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILanguageDirectiveContext)
}

func (s *PackageDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PackageDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitPackageDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) PackageDeclaration() (localctx IPackageDeclarationContext) {
	localctx = NewPackageDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZserioParserRULE_packageDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserPACKAGE {
		{
			p.SetState(140)
			p.PackageNameDefinition()
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserIMPORT {
		{
			p.SetState(143)
			p.ImportDeclaration()
		}


		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 35)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 35))) & ((1 << (ZserioParserBITMASK - 35)) | (1 << (ZserioParserCHOICE - 35)) | (1 << (ZserioParserCONST - 35)) | (1 << (ZserioParserENUM - 35)) | (1 << (ZserioParserINSTANTIATE - 35)) | (1 << (ZserioParserPUBSUB - 35)))) != 0) || ((((_la - 67)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 67))) & ((1 << (ZserioParserRULE_GROUP - 67)) | (1 << (ZserioParserSERVICE - 67)) | (1 << (ZserioParserSQL_DATABASE - 67)) | (1 << (ZserioParserSQL_TABLE - 67)) | (1 << (ZserioParserSTRUCTURE - 67)) | (1 << (ZserioParserSUBTYPE - 67)) | (1 << (ZserioParserUNION - 67)))) != 0) {
		{
			p.SetState(149)
			p.LanguageDirective()
		}


		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(ZserioParserEOF)
	}



	return localctx
}


// IPackageNameDefinitionContext is an interface to support dynamic dispatch.
type IPackageNameDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageNameDefinitionContext differentiates from other interfaces.
	IsPackageNameDefinitionContext()
}

type PackageNameDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageNameDefinitionContext() *PackageNameDefinitionContext {
	var p = new(PackageNameDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_packageNameDefinition
	return p
}

func (*PackageNameDefinitionContext) IsPackageNameDefinitionContext() {}

func NewPackageNameDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageNameDefinitionContext {
	var p = new(PackageNameDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_packageNameDefinition

	return p
}

func (s *PackageNameDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageNameDefinitionContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ZserioParserPACKAGE, 0)
}

func (s *PackageNameDefinitionContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *PackageNameDefinitionContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PackageNameDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *PackageNameDefinitionContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserDOT)
}

func (s *PackageNameDefinitionContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserDOT, i)
}

func (s *PackageNameDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageNameDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PackageNameDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterPackageNameDefinition(s)
	}
}

func (s *PackageNameDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitPackageNameDefinition(s)
	}
}

func (s *PackageNameDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitPackageNameDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) PackageNameDefinition() (localctx IPackageNameDefinitionContext) {
	localctx = NewPackageNameDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZserioParserRULE_packageNameDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ZserioParserPACKAGE)
	}
	{
		p.SetState(158)
		p.Id()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserDOT {
		{
			p.SetState(159)
			p.Match(ZserioParserDOT)
		}
		{
			p.SetState(160)
			p.Id()
		}


		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ZserioParserIMPORT, 0)
}

func (s *ImportDeclarationContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *ImportDeclarationContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ImportDeclarationContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserDOT)
}

func (s *ImportDeclarationContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserDOT, i)
}

func (s *ImportDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ImportDeclarationContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(ZserioParserMULTIPLY, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZserioParserRULE_importDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(ZserioParserIMPORT)
	}
	{
		p.SetState(169)
		p.Id()
	}
	{
		p.SetState(170)
		p.Match(ZserioParserDOT)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(171)
				p.Id()
			}
			{
				p.SetState(172)
				p.Match(ZserioParserDOT)
			}


		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserID:
		{
			p.SetState(179)
			p.Id()
		}


	case ZserioParserMULTIPLY:
		{
			p.SetState(180)
			p.Match(ZserioParserMULTIPLY)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(183)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ILanguageDirectiveContext is an interface to support dynamic dispatch.
type ILanguageDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageDirectiveContext differentiates from other interfaces.
	IsLanguageDirectiveContext()
}

type LanguageDirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageDirectiveContext() *LanguageDirectiveContext {
	var p = new(LanguageDirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_languageDirective
	return p
}

func (*LanguageDirectiveContext) IsLanguageDirectiveContext() {}

func NewLanguageDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageDirectiveContext {
	var p = new(LanguageDirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_languageDirective

	return p
}

func (s *LanguageDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageDirectiveContext) SymbolDefinition() ISymbolDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolDefinitionContext)
}

func (s *LanguageDirectiveContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *LanguageDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LanguageDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLanguageDirective(s)
	}
}

func (s *LanguageDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLanguageDirective(s)
	}
}

func (s *LanguageDirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLanguageDirective(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) LanguageDirective() (localctx ILanguageDirectiveContext) {
	localctx = NewLanguageDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZserioParserRULE_languageDirective)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserCONST, ZserioParserRULE_GROUP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.SymbolDefinition()
		}


	case ZserioParserBITMASK, ZserioParserCHOICE, ZserioParserENUM, ZserioParserINSTANTIATE, ZserioParserPUBSUB, ZserioParserSERVICE, ZserioParserSQL_DATABASE, ZserioParserSQL_TABLE, ZserioParserSTRUCTURE, ZserioParserSUBTYPE, ZserioParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.TypeDeclaration()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) SubtypeDeclaration() ISubtypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtypeDeclarationContext)
}

func (s *TypeDeclarationContext) StructureDeclaration() IStructureDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructureDeclarationContext)
}

func (s *TypeDeclarationContext) ChoiceDeclaration() IChoiceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceDeclarationContext)
}

func (s *TypeDeclarationContext) UnionDeclaration() IUnionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionDeclarationContext)
}

func (s *TypeDeclarationContext) EnumDeclaration() IEnumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclarationContext)
}

func (s *TypeDeclarationContext) BitmaskDeclaration() IBitmaskDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitmaskDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBitmaskDeclarationContext)
}

func (s *TypeDeclarationContext) SqlTableDeclaration() ISqlTableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlTableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlTableDeclarationContext)
}

func (s *TypeDeclarationContext) SqlDatabaseDefinition() ISqlDatabaseDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlDatabaseDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlDatabaseDefinitionContext)
}

func (s *TypeDeclarationContext) ServiceDefinition() IServiceDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceDefinitionContext)
}

func (s *TypeDeclarationContext) PubsubDefinition() IPubsubDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPubsubDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPubsubDefinitionContext)
}

func (s *TypeDeclarationContext) InstantiateDeclaration() IInstantiateDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstantiateDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstantiateDeclarationContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZserioParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserSUBTYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.SubtypeDeclaration()
		}


	case ZserioParserSTRUCTURE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.StructureDeclaration()
		}


	case ZserioParserCHOICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.ChoiceDeclaration()
		}


	case ZserioParserUNION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
			p.UnionDeclaration()
		}


	case ZserioParserENUM:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.EnumDeclaration()
		}


	case ZserioParserBITMASK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(194)
			p.BitmaskDeclaration()
		}


	case ZserioParserSQL_TABLE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(195)
			p.SqlTableDeclaration()
		}


	case ZserioParserSQL_DATABASE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(196)
			p.SqlDatabaseDefinition()
		}


	case ZserioParserSERVICE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(197)
			p.ServiceDefinition()
		}


	case ZserioParserPUBSUB:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(198)
			p.PubsubDefinition()
		}


	case ZserioParserINSTANTIATE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(199)
			p.InstantiateDeclaration()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ISymbolDefinitionContext is an interface to support dynamic dispatch.
type ISymbolDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolDefinitionContext differentiates from other interfaces.
	IsSymbolDefinitionContext()
}

type SymbolDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolDefinitionContext() *SymbolDefinitionContext {
	var p = new(SymbolDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_symbolDefinition
	return p
}

func (*SymbolDefinitionContext) IsSymbolDefinitionContext() {}

func NewSymbolDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolDefinitionContext {
	var p = new(SymbolDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_symbolDefinition

	return p
}

func (s *SymbolDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolDefinitionContext) ConstDefinition() IConstDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDefinitionContext)
}

func (s *SymbolDefinitionContext) RuleGroupDefinition() IRuleGroupDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleGroupDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleGroupDefinitionContext)
}

func (s *SymbolDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SymbolDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSymbolDefinition(s)
	}
}

func (s *SymbolDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSymbolDefinition(s)
	}
}

func (s *SymbolDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSymbolDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SymbolDefinition() (localctx ISymbolDefinitionContext) {
	localctx = NewSymbolDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ZserioParserRULE_symbolDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.ConstDefinition()
		}


	case ZserioParserRULE_GROUP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.RuleGroupDefinition()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IConstDefinitionContext is an interface to support dynamic dispatch.
type IConstDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstDefinitionContext differentiates from other interfaces.
	IsConstDefinitionContext()
}

type ConstDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDefinitionContext() *ConstDefinitionContext {
	var p = new(ConstDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_constDefinition
	return p
}

func (*ConstDefinitionContext) IsConstDefinitionContext() {}

func NewConstDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDefinitionContext {
	var p = new(ConstDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_constDefinition

	return p
}

func (s *ConstDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDefinitionContext) CONST() antlr.TerminalNode {
	return s.GetToken(ZserioParserCONST, 0)
}

func (s *ConstDefinitionContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *ConstDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ConstDefinitionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZserioParserASSIGN, 0)
}

func (s *ConstDefinitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ConstDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterConstDefinition(s)
	}
}

func (s *ConstDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitConstDefinition(s)
	}
}

func (s *ConstDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitConstDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ConstDefinition() (localctx IConstDefinitionContext) {
	localctx = NewConstDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ZserioParserRULE_constDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(ZserioParserCONST)
	}
	{
		p.SetState(207)
		p.TypeInstantiation()
	}
	{
		p.SetState(208)
		p.Id()
	}
	{
		p.SetState(209)
		p.Match(ZserioParserASSIGN)
	}
	{
		p.SetState(210)
		p.expression(0)
	}
	{
		p.SetState(211)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IRuleGroupDefinitionContext is an interface to support dynamic dispatch.
type IRuleGroupDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleGroupDefinitionContext differentiates from other interfaces.
	IsRuleGroupDefinitionContext()
}

type RuleGroupDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleGroupDefinitionContext() *RuleGroupDefinitionContext {
	var p = new(RuleGroupDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_ruleGroupDefinition
	return p
}

func (*RuleGroupDefinitionContext) IsRuleGroupDefinitionContext() {}

func NewRuleGroupDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleGroupDefinitionContext {
	var p = new(RuleGroupDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_ruleGroupDefinition

	return p
}

func (s *RuleGroupDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleGroupDefinitionContext) RULE_GROUP() antlr.TerminalNode {
	return s.GetToken(ZserioParserRULE_GROUP, 0)
}

func (s *RuleGroupDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *RuleGroupDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *RuleGroupDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *RuleGroupDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *RuleGroupDefinitionContext) AllRuleDefinition() []IRuleDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleDefinitionContext)(nil)).Elem())
	var tst = make([]IRuleDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleDefinitionContext)
		}
	}

	return tst
}

func (s *RuleGroupDefinitionContext) RuleDefinition(i int) IRuleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleDefinitionContext)
}

func (s *RuleGroupDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleGroupDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleGroupDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterRuleGroupDefinition(s)
	}
}

func (s *RuleGroupDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitRuleGroupDefinition(s)
	}
}

func (s *RuleGroupDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitRuleGroupDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) RuleGroupDefinition() (localctx IRuleGroupDefinitionContext) {
	localctx = NewRuleGroupDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ZserioParserRULE_ruleGroupDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(ZserioParserRULE_GROUP)
	}
	{
		p.SetState(214)
		p.Id()
	}
	{
		p.SetState(215)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserRULE {
		{
			p.SetState(216)
			p.RuleDefinition()
		}


		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(223)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IRuleDefinitionContext is an interface to support dynamic dispatch.
type IRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDefinitionContext differentiates from other interfaces.
	IsRuleDefinitionContext()
}

type RuleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDefinitionContext() *RuleDefinitionContext {
	var p = new(RuleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_ruleDefinition
	return p
}

func (*RuleDefinitionContext) IsRuleDefinitionContext() {}

func NewRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDefinitionContext {
	var p = new(RuleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_ruleDefinition

	return p
}

func (s *RuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDefinitionContext) RULE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRULE, 0)
}

func (s *RuleDefinitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RuleDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *RuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterRuleDefinition(s)
	}
}

func (s *RuleDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitRuleDefinition(s)
	}
}

func (s *RuleDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitRuleDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) RuleDefinition() (localctx IRuleDefinitionContext) {
	localctx = NewRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ZserioParserRULE_ruleDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(ZserioParserRULE)
	}
	{
		p.SetState(226)
		p.expression(0)
	}
	{
		p.SetState(227)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISubtypeDeclarationContext is an interface to support dynamic dispatch.
type ISubtypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtypeDeclarationContext differentiates from other interfaces.
	IsSubtypeDeclarationContext()
}

type SubtypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtypeDeclarationContext() *SubtypeDeclarationContext {
	var p = new(SubtypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_subtypeDeclaration
	return p
}

func (*SubtypeDeclarationContext) IsSubtypeDeclarationContext() {}

func NewSubtypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtypeDeclarationContext {
	var p = new(SubtypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_subtypeDeclaration

	return p
}

func (s *SubtypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtypeDeclarationContext) SUBTYPE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSUBTYPE, 0)
}

func (s *SubtypeDeclarationContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *SubtypeDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SubtypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SubtypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SubtypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSubtypeDeclaration(s)
	}
}

func (s *SubtypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSubtypeDeclaration(s)
	}
}

func (s *SubtypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSubtypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SubtypeDeclaration() (localctx ISubtypeDeclarationContext) {
	localctx = NewSubtypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ZserioParserRULE_subtypeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(ZserioParserSUBTYPE)
	}
	{
		p.SetState(230)
		p.TypeReference()
	}
	{
		p.SetState(231)
		p.Id()
	}
	{
		p.SetState(232)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IStructureDeclarationContext is an interface to support dynamic dispatch.
type IStructureDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureDeclarationContext differentiates from other interfaces.
	IsStructureDeclarationContext()
}

type StructureDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureDeclarationContext() *StructureDeclarationContext {
	var p = new(StructureDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_structureDeclaration
	return p
}

func (*StructureDeclarationContext) IsStructureDeclarationContext() {}

func NewStructureDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureDeclarationContext {
	var p = new(StructureDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_structureDeclaration

	return p
}

func (s *StructureDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureDeclarationContext) STRUCTURE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSTRUCTURE, 0)
}

func (s *StructureDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *StructureDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *StructureDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *StructureDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *StructureDeclarationContext) TemplateParameters() ITemplateParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersContext)
}

func (s *StructureDeclarationContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *StructureDeclarationContext) AllStructureFieldDefinition() []IStructureFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructureFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IStructureFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructureFieldDefinitionContext)
		}
	}

	return tst
}

func (s *StructureDeclarationContext) StructureFieldDefinition(i int) IStructureFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructureFieldDefinitionContext)
}

func (s *StructureDeclarationContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem())
	var tst = make([]IFunctionDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDefinitionContext)
		}
	}

	return tst
}

func (s *StructureDeclarationContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *StructureDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructureDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterStructureDeclaration(s)
	}
}

func (s *StructureDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitStructureDeclaration(s)
	}
}

func (s *StructureDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitStructureDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) StructureDeclaration() (localctx IStructureDeclarationContext) {
	localctx = NewStructureDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ZserioParserRULE_structureDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(ZserioParserSTRUCTURE)
	}
	{
		p.SetState(235)
		p.Id()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLT {
		{
			p.SetState(236)
			p.TemplateParameters()
		}

	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLPAREN {
		{
			p.SetState(239)
			p.TypeParameters()
		}

	}
	{
		p.SetState(242)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZserioParserBANG) | (1 << ZserioParserLPAREN) | (1 << ZserioParserMINUS) | (1 << ZserioParserPLUS) | (1 << ZserioParserTILDE))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (ZserioParserALIGN - 32)) | (1 << (ZserioParserBIT_FIELD - 32)) | (1 << (ZserioParserBOOL - 32)) | (1 << (ZserioParserEXTERN - 32)) | (1 << (ZserioParserFLOAT16 - 32)) | (1 << (ZserioParserFLOAT32 - 32)) | (1 << (ZserioParserFLOAT64 - 32)) | (1 << (ZserioParserIMPLICIT - 32)) | (1 << (ZserioParserINDEX - 32)) | (1 << (ZserioParserINT_FIELD - 32)) | (1 << (ZserioParserINT16 - 32)) | (1 << (ZserioParserINT32 - 32)) | (1 << (ZserioParserINT64 - 32)) | (1 << (ZserioParserINT8 - 32)) | (1 << (ZserioParserLENGTHOF - 32)) | (1 << (ZserioParserNUMBITS - 32)) | (1 << (ZserioParserOPTIONAL - 32)) | (1 << (ZserioParserPACKED - 32)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVALUEOF - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)) | (1 << (ZserioParserBOOL_LITERAL - 74)) | (1 << (ZserioParserSTRING_LITERAL - 74)) | (1 << (ZserioParserBINARY_LITERAL - 74)) | (1 << (ZserioParserOCTAL_LITERAL - 74)) | (1 << (ZserioParserHEXADECIMAL_LITERAL - 74)) | (1 << (ZserioParserDOUBLE_LITERAL - 74)))) != 0) || ((((_la - 106)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 106))) & ((1 << (ZserioParserFLOAT_LITERAL - 106)) | (1 << (ZserioParserDECIMAL_LITERAL - 106)) | (1 << (ZserioParserID - 106)))) != 0) {
		{
			p.SetState(243)
			p.StructureFieldDefinition()
		}


		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserFUNCTION {
		{
			p.SetState(249)
			p.FunctionDefinition()
		}


		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(255)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(256)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IStructureFieldDefinitionContext is an interface to support dynamic dispatch.
type IStructureFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureFieldDefinitionContext differentiates from other interfaces.
	IsStructureFieldDefinitionContext()
}

type StructureFieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureFieldDefinitionContext() *StructureFieldDefinitionContext {
	var p = new(StructureFieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_structureFieldDefinition
	return p
}

func (*StructureFieldDefinitionContext) IsStructureFieldDefinitionContext() {}

func NewStructureFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureFieldDefinitionContext {
	var p = new(StructureFieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_structureFieldDefinition

	return p
}

func (s *StructureFieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureFieldDefinitionContext) FieldTypeId() IFieldTypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeIdContext)
}

func (s *StructureFieldDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *StructureFieldDefinitionContext) FieldAlignment() IFieldAlignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldAlignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldAlignmentContext)
}

func (s *StructureFieldDefinitionContext) FieldOffset() IFieldOffsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOffsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOffsetContext)
}

func (s *StructureFieldDefinitionContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserOPTIONAL, 0)
}

func (s *StructureFieldDefinitionContext) FieldInitializer() IFieldInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldInitializerContext)
}

func (s *StructureFieldDefinitionContext) FieldOptionalClause() IFieldOptionalClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionalClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionalClauseContext)
}

func (s *StructureFieldDefinitionContext) FieldConstraint() IFieldConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldConstraintContext)
}

func (s *StructureFieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureFieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructureFieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterStructureFieldDefinition(s)
	}
}

func (s *StructureFieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitStructureFieldDefinition(s)
	}
}

func (s *StructureFieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitStructureFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) StructureFieldDefinition() (localctx IStructureFieldDefinitionContext) {
	localctx = NewStructureFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZserioParserRULE_structureFieldDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserALIGN {
		{
			p.SetState(258)
			p.FieldAlignment()
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(261)
			p.FieldOffset()
		}


	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserOPTIONAL {
		{
			p.SetState(264)
			p.Match(ZserioParserOPTIONAL)
		}

	}
	{
		p.SetState(267)
		p.FieldTypeId()
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserASSIGN {
		{
			p.SetState(268)
			p.FieldInitializer()
		}

	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserIF {
		{
			p.SetState(271)
			p.FieldOptionalClause()
		}

	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserCOLON {
		{
			p.SetState(274)
			p.FieldConstraint()
		}

	}
	{
		p.SetState(277)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IFieldAlignmentContext is an interface to support dynamic dispatch.
type IFieldAlignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldAlignmentContext differentiates from other interfaces.
	IsFieldAlignmentContext()
}

type FieldAlignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldAlignmentContext() *FieldAlignmentContext {
	var p = new(FieldAlignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldAlignment
	return p
}

func (*FieldAlignmentContext) IsFieldAlignmentContext() {}

func NewFieldAlignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAlignmentContext {
	var p = new(FieldAlignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldAlignment

	return p
}

func (s *FieldAlignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAlignmentContext) ALIGN() antlr.TerminalNode {
	return s.GetToken(ZserioParserALIGN, 0)
}

func (s *FieldAlignmentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *FieldAlignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldAlignmentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *FieldAlignmentContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *FieldAlignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAlignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldAlignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldAlignment(s)
	}
}

func (s *FieldAlignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldAlignment(s)
	}
}

func (s *FieldAlignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldAlignment(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldAlignment() (localctx IFieldAlignmentContext) {
	localctx = NewFieldAlignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ZserioParserRULE_fieldAlignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(ZserioParserALIGN)
	}
	{
		p.SetState(280)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(281)
		p.expression(0)
	}
	{
		p.SetState(282)
		p.Match(ZserioParserRPAREN)
	}
	{
		p.SetState(283)
		p.Match(ZserioParserCOLON)
	}



	return localctx
}


// IFieldOffsetContext is an interface to support dynamic dispatch.
type IFieldOffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOffsetContext differentiates from other interfaces.
	IsFieldOffsetContext()
}

type FieldOffsetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOffsetContext() *FieldOffsetContext {
	var p = new(FieldOffsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldOffset
	return p
}

func (*FieldOffsetContext) IsFieldOffsetContext() {}

func NewFieldOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOffsetContext {
	var p = new(FieldOffsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldOffset

	return p
}

func (s *FieldOffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOffsetContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldOffsetContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *FieldOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldOffset(s)
	}
}

func (s *FieldOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldOffset(s)
	}
}

func (s *FieldOffsetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldOffset(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldOffset() (localctx IFieldOffsetContext) {
	localctx = NewFieldOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZserioParserRULE_fieldOffset)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.expression(0)
	}
	{
		p.SetState(286)
		p.Match(ZserioParserCOLON)
	}



	return localctx
}


// IFieldTypeIdContext is an interface to support dynamic dispatch.
type IFieldTypeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeIdContext differentiates from other interfaces.
	IsFieldTypeIdContext()
}

type FieldTypeIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeIdContext() *FieldTypeIdContext {
	var p = new(FieldTypeIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldTypeId
	return p
}

func (*FieldTypeIdContext) IsFieldTypeIdContext() {}

func NewFieldTypeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeIdContext {
	var p = new(FieldTypeIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldTypeId

	return p
}

func (s *FieldTypeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeIdContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *FieldTypeIdContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *FieldTypeIdContext) PACKED() antlr.TerminalNode {
	return s.GetToken(ZserioParserPACKED, 0)
}

func (s *FieldTypeIdContext) IMPLICIT() antlr.TerminalNode {
	return s.GetToken(ZserioParserIMPLICIT, 0)
}

func (s *FieldTypeIdContext) FieldArrayRange() IFieldArrayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldArrayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldArrayRangeContext)
}

func (s *FieldTypeIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldTypeIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldTypeId(s)
	}
}

func (s *FieldTypeIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldTypeId(s)
	}
}

func (s *FieldTypeIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldTypeId(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldTypeId() (localctx IFieldTypeIdContext) {
	localctx = NewFieldTypeIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ZserioParserRULE_fieldTypeId)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserPACKED {
		{
			p.SetState(288)
			p.Match(ZserioParserPACKED)
		}

	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserIMPLICIT {
		{
			p.SetState(291)
			p.Match(ZserioParserIMPLICIT)
		}

	}
	{
		p.SetState(294)
		p.TypeInstantiation()
	}
	{
		p.SetState(295)
		p.Id()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLBRACKET {
		{
			p.SetState(296)
			p.FieldArrayRange()
		}

	}



	return localctx
}


// IFieldArrayRangeContext is an interface to support dynamic dispatch.
type IFieldArrayRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldArrayRangeContext differentiates from other interfaces.
	IsFieldArrayRangeContext()
}

type FieldArrayRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldArrayRangeContext() *FieldArrayRangeContext {
	var p = new(FieldArrayRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldArrayRange
	return p
}

func (*FieldArrayRangeContext) IsFieldArrayRangeContext() {}

func NewFieldArrayRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldArrayRangeContext {
	var p = new(FieldArrayRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldArrayRange

	return p
}

func (s *FieldArrayRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldArrayRangeContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACKET, 0)
}

func (s *FieldArrayRangeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACKET, 0)
}

func (s *FieldArrayRangeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldArrayRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldArrayRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldArrayRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldArrayRange(s)
	}
}

func (s *FieldArrayRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldArrayRange(s)
	}
}

func (s *FieldArrayRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldArrayRange(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldArrayRange() (localctx IFieldArrayRangeContext) {
	localctx = NewFieldArrayRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ZserioParserRULE_fieldArrayRange)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(ZserioParserLBRACKET)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZserioParserBANG) | (1 << ZserioParserLPAREN) | (1 << ZserioParserMINUS) | (1 << ZserioParserPLUS) | (1 << ZserioParserTILDE))) != 0) || ((((_la - 50)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 50))) & ((1 << (ZserioParserINDEX - 50)) | (1 << (ZserioParserLENGTHOF - 50)) | (1 << (ZserioParserNUMBITS - 50)))) != 0) || ((((_la - 85)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 85))) & ((1 << (ZserioParserVALUEOF - 85)) | (1 << (ZserioParserBOOL_LITERAL - 85)) | (1 << (ZserioParserSTRING_LITERAL - 85)) | (1 << (ZserioParserBINARY_LITERAL - 85)) | (1 << (ZserioParserOCTAL_LITERAL - 85)) | (1 << (ZserioParserHEXADECIMAL_LITERAL - 85)) | (1 << (ZserioParserDOUBLE_LITERAL - 85)) | (1 << (ZserioParserFLOAT_LITERAL - 85)) | (1 << (ZserioParserDECIMAL_LITERAL - 85)) | (1 << (ZserioParserID - 85)))) != 0) {
		{
			p.SetState(300)
			p.expression(0)
		}

	}
	{
		p.SetState(303)
		p.Match(ZserioParserRBRACKET)
	}



	return localctx
}


// IFieldInitializerContext is an interface to support dynamic dispatch.
type IFieldInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldInitializerContext differentiates from other interfaces.
	IsFieldInitializerContext()
}

type FieldInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldInitializerContext() *FieldInitializerContext {
	var p = new(FieldInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldInitializer
	return p
}

func (*FieldInitializerContext) IsFieldInitializerContext() {}

func NewFieldInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitializerContext {
	var p = new(FieldInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldInitializer

	return p
}

func (s *FieldInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldInitializerContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZserioParserASSIGN, 0)
}

func (s *FieldInitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldInitializer(s)
	}
}

func (s *FieldInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldInitializer(s)
	}
}

func (s *FieldInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldInitializer() (localctx IFieldInitializerContext) {
	localctx = NewFieldInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ZserioParserRULE_fieldInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(ZserioParserASSIGN)
	}
	{
		p.SetState(306)
		p.expression(0)
	}



	return localctx
}


// IFieldOptionalClauseContext is an interface to support dynamic dispatch.
type IFieldOptionalClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionalClauseContext differentiates from other interfaces.
	IsFieldOptionalClauseContext()
}

type FieldOptionalClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionalClauseContext() *FieldOptionalClauseContext {
	var p = new(FieldOptionalClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldOptionalClause
	return p
}

func (*FieldOptionalClauseContext) IsFieldOptionalClauseContext() {}

func NewFieldOptionalClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionalClauseContext {
	var p = new(FieldOptionalClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldOptionalClause

	return p
}

func (s *FieldOptionalClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionalClauseContext) IF() antlr.TerminalNode {
	return s.GetToken(ZserioParserIF, 0)
}

func (s *FieldOptionalClauseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldOptionalClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionalClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldOptionalClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldOptionalClause(s)
	}
}

func (s *FieldOptionalClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldOptionalClause(s)
	}
}

func (s *FieldOptionalClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldOptionalClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldOptionalClause() (localctx IFieldOptionalClauseContext) {
	localctx = NewFieldOptionalClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ZserioParserRULE_fieldOptionalClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(ZserioParserIF)
	}
	{
		p.SetState(309)
		p.expression(0)
	}



	return localctx
}


// IFieldConstraintContext is an interface to support dynamic dispatch.
type IFieldConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldConstraintContext differentiates from other interfaces.
	IsFieldConstraintContext()
}

type FieldConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldConstraintContext() *FieldConstraintContext {
	var p = new(FieldConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fieldConstraint
	return p
}

func (*FieldConstraintContext) IsFieldConstraintContext() {}

func NewFieldConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldConstraintContext {
	var p = new(FieldConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fieldConstraint

	return p
}

func (s *FieldConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldConstraintContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *FieldConstraintContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFieldConstraint(s)
	}
}

func (s *FieldConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFieldConstraint(s)
	}
}

func (s *FieldConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFieldConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FieldConstraint() (localctx IFieldConstraintContext) {
	localctx = NewFieldConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ZserioParserRULE_fieldConstraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(ZserioParserCOLON)
	}
	{
		p.SetState(312)
		p.expression(0)
	}



	return localctx
}


// IChoiceDeclarationContext is an interface to support dynamic dispatch.
type IChoiceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceDeclarationContext differentiates from other interfaces.
	IsChoiceDeclarationContext()
}

type ChoiceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceDeclarationContext() *ChoiceDeclarationContext {
	var p = new(ChoiceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_choiceDeclaration
	return p
}

func (*ChoiceDeclarationContext) IsChoiceDeclarationContext() {}

func NewChoiceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceDeclarationContext {
	var p = new(ChoiceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_choiceDeclaration

	return p
}

func (s *ChoiceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceDeclarationContext) CHOICE() antlr.TerminalNode {
	return s.GetToken(ZserioParserCHOICE, 0)
}

func (s *ChoiceDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ChoiceDeclarationContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *ChoiceDeclarationContext) ON() antlr.TerminalNode {
	return s.GetToken(ZserioParserON, 0)
}

func (s *ChoiceDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ChoiceDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *ChoiceDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *ChoiceDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ChoiceDeclarationContext) TemplateParameters() ITemplateParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersContext)
}

func (s *ChoiceDeclarationContext) AllChoiceCases() []IChoiceCasesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChoiceCasesContext)(nil)).Elem())
	var tst = make([]IChoiceCasesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChoiceCasesContext)
		}
	}

	return tst
}

func (s *ChoiceDeclarationContext) ChoiceCases(i int) IChoiceCasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceCasesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChoiceCasesContext)
}

func (s *ChoiceDeclarationContext) ChoiceDefault() IChoiceDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceDefaultContext)
}

func (s *ChoiceDeclarationContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem())
	var tst = make([]IFunctionDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDefinitionContext)
		}
	}

	return tst
}

func (s *ChoiceDeclarationContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *ChoiceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ChoiceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterChoiceDeclaration(s)
	}
}

func (s *ChoiceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitChoiceDeclaration(s)
	}
}

func (s *ChoiceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitChoiceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ChoiceDeclaration() (localctx IChoiceDeclarationContext) {
	localctx = NewChoiceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ZserioParserRULE_choiceDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(ZserioParserCHOICE)
	}
	{
		p.SetState(315)
		p.Id()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLT {
		{
			p.SetState(316)
			p.TemplateParameters()
		}

	}
	{
		p.SetState(319)
		p.TypeParameters()
	}
	{
		p.SetState(320)
		p.Match(ZserioParserON)
	}
	{
		p.SetState(321)
		p.expression(0)
	}
	{
		p.SetState(322)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserCASE {
		{
			p.SetState(323)
			p.ChoiceCases()
		}


		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserDEFAULT {
		{
			p.SetState(329)
			p.ChoiceDefault()
		}

	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserFUNCTION {
		{
			p.SetState(332)
			p.FunctionDefinition()
		}


		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(338)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(339)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IChoiceCasesContext is an interface to support dynamic dispatch.
type IChoiceCasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceCasesContext differentiates from other interfaces.
	IsChoiceCasesContext()
}

type ChoiceCasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceCasesContext() *ChoiceCasesContext {
	var p = new(ChoiceCasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_choiceCases
	return p
}

func (*ChoiceCasesContext) IsChoiceCasesContext() {}

func NewChoiceCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceCasesContext {
	var p = new(ChoiceCasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_choiceCases

	return p
}

func (s *ChoiceCasesContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceCasesContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ChoiceCasesContext) AllChoiceCase() []IChoiceCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChoiceCaseContext)(nil)).Elem())
	var tst = make([]IChoiceCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChoiceCaseContext)
		}
	}

	return tst
}

func (s *ChoiceCasesContext) ChoiceCase(i int) IChoiceCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChoiceCaseContext)
}

func (s *ChoiceCasesContext) ChoiceFieldDefinition() IChoiceFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceFieldDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceFieldDefinitionContext)
}

func (s *ChoiceCasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceCasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ChoiceCasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterChoiceCases(s)
	}
}

func (s *ChoiceCasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitChoiceCases(s)
	}
}

func (s *ChoiceCasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitChoiceCases(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ChoiceCases() (localctx IChoiceCasesContext) {
	localctx = NewChoiceCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ZserioParserRULE_choiceCases)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == ZserioParserCASE {
		{
			p.SetState(341)
			p.ChoiceCase()
		}


		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserIMPLICIT - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)) | (1 << (ZserioParserPACKED - 33)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(346)
			p.ChoiceFieldDefinition()
		}

	}
	{
		p.SetState(349)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IChoiceCaseContext is an interface to support dynamic dispatch.
type IChoiceCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceCaseContext differentiates from other interfaces.
	IsChoiceCaseContext()
}

type ChoiceCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceCaseContext() *ChoiceCaseContext {
	var p = new(ChoiceCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_choiceCase
	return p
}

func (*ChoiceCaseContext) IsChoiceCaseContext() {}

func NewChoiceCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceCaseContext {
	var p = new(ChoiceCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_choiceCase

	return p
}

func (s *ChoiceCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(ZserioParserCASE, 0)
}

func (s *ChoiceCaseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ChoiceCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *ChoiceCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ChoiceCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterChoiceCase(s)
	}
}

func (s *ChoiceCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitChoiceCase(s)
	}
}

func (s *ChoiceCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitChoiceCase(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ChoiceCase() (localctx IChoiceCaseContext) {
	localctx = NewChoiceCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ZserioParserRULE_choiceCase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(ZserioParserCASE)
	}
	{
		p.SetState(352)
		p.expression(0)
	}
	{
		p.SetState(353)
		p.Match(ZserioParserCOLON)
	}



	return localctx
}


// IChoiceDefaultContext is an interface to support dynamic dispatch.
type IChoiceDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceDefaultContext differentiates from other interfaces.
	IsChoiceDefaultContext()
}

type ChoiceDefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceDefaultContext() *ChoiceDefaultContext {
	var p = new(ChoiceDefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_choiceDefault
	return p
}

func (*ChoiceDefaultContext) IsChoiceDefaultContext() {}

func NewChoiceDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceDefaultContext {
	var p = new(ChoiceDefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_choiceDefault

	return p
}

func (s *ChoiceDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceDefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ZserioParserDEFAULT, 0)
}

func (s *ChoiceDefaultContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *ChoiceDefaultContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ChoiceDefaultContext) ChoiceFieldDefinition() IChoiceFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceFieldDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceFieldDefinitionContext)
}

func (s *ChoiceDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ChoiceDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterChoiceDefault(s)
	}
}

func (s *ChoiceDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitChoiceDefault(s)
	}
}

func (s *ChoiceDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitChoiceDefault(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ChoiceDefault() (localctx IChoiceDefaultContext) {
	localctx = NewChoiceDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ZserioParserRULE_choiceDefault)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(ZserioParserDEFAULT)
	}
	{
		p.SetState(356)
		p.Match(ZserioParserCOLON)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserIMPLICIT - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)) | (1 << (ZserioParserPACKED - 33)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(357)
			p.ChoiceFieldDefinition()
		}

	}
	{
		p.SetState(360)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IChoiceFieldDefinitionContext is an interface to support dynamic dispatch.
type IChoiceFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChoiceFieldDefinitionContext differentiates from other interfaces.
	IsChoiceFieldDefinitionContext()
}

type ChoiceFieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChoiceFieldDefinitionContext() *ChoiceFieldDefinitionContext {
	var p = new(ChoiceFieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_choiceFieldDefinition
	return p
}

func (*ChoiceFieldDefinitionContext) IsChoiceFieldDefinitionContext() {}

func NewChoiceFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChoiceFieldDefinitionContext {
	var p = new(ChoiceFieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_choiceFieldDefinition

	return p
}

func (s *ChoiceFieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ChoiceFieldDefinitionContext) FieldTypeId() IFieldTypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeIdContext)
}

func (s *ChoiceFieldDefinitionContext) FieldConstraint() IFieldConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldConstraintContext)
}

func (s *ChoiceFieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChoiceFieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ChoiceFieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterChoiceFieldDefinition(s)
	}
}

func (s *ChoiceFieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitChoiceFieldDefinition(s)
	}
}

func (s *ChoiceFieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitChoiceFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ChoiceFieldDefinition() (localctx IChoiceFieldDefinitionContext) {
	localctx = NewChoiceFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ZserioParserRULE_choiceFieldDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.FieldTypeId()
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserCOLON {
		{
			p.SetState(363)
			p.FieldConstraint()
		}

	}



	return localctx
}


// IUnionDeclarationContext is an interface to support dynamic dispatch.
type IUnionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionDeclarationContext differentiates from other interfaces.
	IsUnionDeclarationContext()
}

type UnionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionDeclarationContext() *UnionDeclarationContext {
	var p = new(UnionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_unionDeclaration
	return p
}

func (*UnionDeclarationContext) IsUnionDeclarationContext() {}

func NewUnionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionDeclarationContext {
	var p = new(UnionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_unionDeclaration

	return p
}

func (s *UnionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionDeclarationContext) UNION() antlr.TerminalNode {
	return s.GetToken(ZserioParserUNION, 0)
}

func (s *UnionDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *UnionDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *UnionDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *UnionDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *UnionDeclarationContext) TemplateParameters() ITemplateParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersContext)
}

func (s *UnionDeclarationContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *UnionDeclarationContext) AllUnionFieldDefinition() []IUnionFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnionFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IUnionFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnionFieldDefinitionContext)
		}
	}

	return tst
}

func (s *UnionDeclarationContext) UnionFieldDefinition(i int) IUnionFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnionFieldDefinitionContext)
}

func (s *UnionDeclarationContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem())
	var tst = make([]IFunctionDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDefinitionContext)
		}
	}

	return tst
}

func (s *UnionDeclarationContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *UnionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterUnionDeclaration(s)
	}
}

func (s *UnionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitUnionDeclaration(s)
	}
}

func (s *UnionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitUnionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) UnionDeclaration() (localctx IUnionDeclarationContext) {
	localctx = NewUnionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ZserioParserRULE_unionDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(ZserioParserUNION)
	}
	{
		p.SetState(367)
		p.Id()
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLT {
		{
			p.SetState(368)
			p.TemplateParameters()
		}

	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLPAREN {
		{
			p.SetState(371)
			p.TypeParameters()
		}

	}
	{
		p.SetState(374)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserIMPLICIT - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)) | (1 << (ZserioParserPACKED - 33)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(375)
			p.UnionFieldDefinition()
		}


		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserFUNCTION {
		{
			p.SetState(381)
			p.FunctionDefinition()
		}


		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(387)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(388)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IUnionFieldDefinitionContext is an interface to support dynamic dispatch.
type IUnionFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionFieldDefinitionContext differentiates from other interfaces.
	IsUnionFieldDefinitionContext()
}

type UnionFieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionFieldDefinitionContext() *UnionFieldDefinitionContext {
	var p = new(UnionFieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_unionFieldDefinition
	return p
}

func (*UnionFieldDefinitionContext) IsUnionFieldDefinitionContext() {}

func NewUnionFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionFieldDefinitionContext {
	var p = new(UnionFieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_unionFieldDefinition

	return p
}

func (s *UnionFieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionFieldDefinitionContext) ChoiceFieldDefinition() IChoiceFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChoiceFieldDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChoiceFieldDefinitionContext)
}

func (s *UnionFieldDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *UnionFieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionFieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnionFieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterUnionFieldDefinition(s)
	}
}

func (s *UnionFieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitUnionFieldDefinition(s)
	}
}

func (s *UnionFieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitUnionFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) UnionFieldDefinition() (localctx IUnionFieldDefinitionContext) {
	localctx = NewUnionFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ZserioParserRULE_unionFieldDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.ChoiceFieldDefinition()
	}
	{
		p.SetState(391)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IEnumDeclarationContext is an interface to support dynamic dispatch.
type IEnumDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDeclarationContext differentiates from other interfaces.
	IsEnumDeclarationContext()
}

type EnumDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclarationContext() *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_enumDeclaration
	return p
}

func (*EnumDeclarationContext) IsEnumDeclarationContext() {}

func NewEnumDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_enumDeclaration

	return p
}

func (s *EnumDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclarationContext) ENUM() antlr.TerminalNode {
	return s.GetToken(ZserioParserENUM, 0)
}

func (s *EnumDeclarationContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *EnumDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *EnumDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *EnumDeclarationContext) AllEnumItem() []IEnumItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumItemContext)(nil)).Elem())
	var tst = make([]IEnumItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumItemContext)
		}
	}

	return tst
}

func (s *EnumDeclarationContext) EnumItem(i int) IEnumItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumItemContext)
}

func (s *EnumDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *EnumDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *EnumDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *EnumDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *EnumDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EnumDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterEnumDeclaration(s)
	}
}

func (s *EnumDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitEnumDeclaration(s)
	}
}

func (s *EnumDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitEnumDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) EnumDeclaration() (localctx IEnumDeclarationContext) {
	localctx = NewEnumDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ZserioParserRULE_enumDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(ZserioParserENUM)
	}
	{
		p.SetState(394)
		p.TypeInstantiation()
	}
	{
		p.SetState(395)
		p.Id()
	}
	{
		p.SetState(396)
		p.Match(ZserioParserLBRACE)
	}
	{
		p.SetState(397)
		p.EnumItem()
	}
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(398)
				p.Match(ZserioParserCOMMA)
			}
			{
				p.SetState(399)
				p.EnumItem()
			}


		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserCOMMA {
		{
			p.SetState(405)
			p.Match(ZserioParserCOMMA)
		}

	}
	{
		p.SetState(408)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(409)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IEnumItemContext is an interface to support dynamic dispatch.
type IEnumItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumItemContext differentiates from other interfaces.
	IsEnumItemContext()
}

type EnumItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumItemContext() *EnumItemContext {
	var p = new(EnumItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_enumItem
	return p
}

func (*EnumItemContext) IsEnumItemContext() {}

func NewEnumItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumItemContext {
	var p = new(EnumItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_enumItem

	return p
}

func (s *EnumItemContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumItemContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *EnumItemContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZserioParserASSIGN, 0)
}

func (s *EnumItemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EnumItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EnumItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterEnumItem(s)
	}
}

func (s *EnumItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitEnumItem(s)
	}
}

func (s *EnumItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitEnumItem(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) EnumItem() (localctx IEnumItemContext) {
	localctx = NewEnumItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ZserioParserRULE_enumItem)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Id()
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserASSIGN {
		{
			p.SetState(412)
			p.Match(ZserioParserASSIGN)
		}
		{
			p.SetState(413)
			p.expression(0)
		}

	}



	return localctx
}


// IBitmaskDeclarationContext is an interface to support dynamic dispatch.
type IBitmaskDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitmaskDeclarationContext differentiates from other interfaces.
	IsBitmaskDeclarationContext()
}

type BitmaskDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitmaskDeclarationContext() *BitmaskDeclarationContext {
	var p = new(BitmaskDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_bitmaskDeclaration
	return p
}

func (*BitmaskDeclarationContext) IsBitmaskDeclarationContext() {}

func NewBitmaskDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitmaskDeclarationContext {
	var p = new(BitmaskDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_bitmaskDeclaration

	return p
}

func (s *BitmaskDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *BitmaskDeclarationContext) BITMASK() antlr.TerminalNode {
	return s.GetToken(ZserioParserBITMASK, 0)
}

func (s *BitmaskDeclarationContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *BitmaskDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *BitmaskDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *BitmaskDeclarationContext) AllBitmaskValue() []IBitmaskValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBitmaskValueContext)(nil)).Elem())
	var tst = make([]IBitmaskValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBitmaskValueContext)
		}
	}

	return tst
}

func (s *BitmaskDeclarationContext) BitmaskValue(i int) IBitmaskValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitmaskValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBitmaskValueContext)
}

func (s *BitmaskDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *BitmaskDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *BitmaskDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *BitmaskDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *BitmaskDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitmaskDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BitmaskDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBitmaskDeclaration(s)
	}
}

func (s *BitmaskDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBitmaskDeclaration(s)
	}
}

func (s *BitmaskDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBitmaskDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) BitmaskDeclaration() (localctx IBitmaskDeclarationContext) {
	localctx = NewBitmaskDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ZserioParserRULE_bitmaskDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Match(ZserioParserBITMASK)
	}
	{
		p.SetState(417)
		p.TypeInstantiation()
	}
	{
		p.SetState(418)
		p.Id()
	}
	{
		p.SetState(419)
		p.Match(ZserioParserLBRACE)
	}
	{
		p.SetState(420)
		p.BitmaskValue()
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(421)
				p.Match(ZserioParserCOMMA)
			}
			{
				p.SetState(422)
				p.BitmaskValue()
			}


		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserCOMMA {
		{
			p.SetState(428)
			p.Match(ZserioParserCOMMA)
		}

	}
	{
		p.SetState(431)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(432)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IBitmaskValueContext is an interface to support dynamic dispatch.
type IBitmaskValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitmaskValueContext differentiates from other interfaces.
	IsBitmaskValueContext()
}

type BitmaskValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitmaskValueContext() *BitmaskValueContext {
	var p = new(BitmaskValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_bitmaskValue
	return p
}

func (*BitmaskValueContext) IsBitmaskValueContext() {}

func NewBitmaskValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitmaskValueContext {
	var p = new(BitmaskValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_bitmaskValue

	return p
}

func (s *BitmaskValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BitmaskValueContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *BitmaskValueContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZserioParserASSIGN, 0)
}

func (s *BitmaskValueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitmaskValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitmaskValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BitmaskValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBitmaskValue(s)
	}
}

func (s *BitmaskValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBitmaskValue(s)
	}
}

func (s *BitmaskValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBitmaskValue(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) BitmaskValue() (localctx IBitmaskValueContext) {
	localctx = NewBitmaskValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ZserioParserRULE_bitmaskValue)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Id()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserASSIGN {
		{
			p.SetState(435)
			p.Match(ZserioParserASSIGN)
		}
		{
			p.SetState(436)
			p.expression(0)
		}

	}



	return localctx
}


// ISqlTableDeclarationContext is an interface to support dynamic dispatch.
type ISqlTableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlTableDeclarationContext differentiates from other interfaces.
	IsSqlTableDeclarationContext()
}

type SqlTableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlTableDeclarationContext() *SqlTableDeclarationContext {
	var p = new(SqlTableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlTableDeclaration
	return p
}

func (*SqlTableDeclarationContext) IsSqlTableDeclarationContext() {}

func NewSqlTableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlTableDeclarationContext {
	var p = new(SqlTableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlTableDeclaration

	return p
}

func (s *SqlTableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlTableDeclarationContext) SQL_TABLE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSQL_TABLE, 0)
}

func (s *SqlTableDeclarationContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *SqlTableDeclarationContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SqlTableDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *SqlTableDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *SqlTableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlTableDeclarationContext) TemplateParameters() ITemplateParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersContext)
}

func (s *SqlTableDeclarationContext) USING() antlr.TerminalNode {
	return s.GetToken(ZserioParserUSING, 0)
}

func (s *SqlTableDeclarationContext) AllSqlTableFieldDefinition() []ISqlTableFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISqlTableFieldDefinitionContext)(nil)).Elem())
	var tst = make([]ISqlTableFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISqlTableFieldDefinitionContext)
		}
	}

	return tst
}

func (s *SqlTableDeclarationContext) SqlTableFieldDefinition(i int) ISqlTableFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlTableFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISqlTableFieldDefinitionContext)
}

func (s *SqlTableDeclarationContext) SqlConstraintDefinition() ISqlConstraintDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlConstraintDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlConstraintDefinitionContext)
}

func (s *SqlTableDeclarationContext) SqlWithoutRowId() ISqlWithoutRowIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlWithoutRowIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlWithoutRowIdContext)
}

func (s *SqlTableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlTableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlTableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlTableDeclaration(s)
	}
}

func (s *SqlTableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlTableDeclaration(s)
	}
}

func (s *SqlTableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlTableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlTableDeclaration() (localctx ISqlTableDeclarationContext) {
	localctx = NewSqlTableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ZserioParserRULE_sqlTableDeclaration)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(ZserioParserSQL_TABLE)
	}
	{
		p.SetState(440)
		p.Id()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserLT {
		{
			p.SetState(441)
			p.TemplateParameters()
		}

	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserUSING {
		{
			p.SetState(444)
			p.Match(ZserioParserUSING)
		}
		{
			p.SetState(445)
			p.Id()
		}

	}
	{
		p.SetState(448)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)))) != 0) || ((((_la - 72)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 72))) & ((1 << (ZserioParserSQL_VIRTUAL - 72)) | (1 << (ZserioParserSTRING - 72)) | (1 << (ZserioParserUINT16 - 72)) | (1 << (ZserioParserUINT32 - 72)) | (1 << (ZserioParserUINT64 - 72)) | (1 << (ZserioParserUINT8 - 72)) | (1 << (ZserioParserVARINT - 72)) | (1 << (ZserioParserVARINT16 - 72)) | (1 << (ZserioParserVARINT32 - 72)) | (1 << (ZserioParserVARINT64 - 72)) | (1 << (ZserioParserVARSIZE - 72)) | (1 << (ZserioParserVARUINT - 72)) | (1 << (ZserioParserVARUINT16 - 72)) | (1 << (ZserioParserVARUINT32 - 72)) | (1 << (ZserioParserVARUINT64 - 72)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(449)
			p.SqlTableFieldDefinition()
		}


		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserSQL {
		{
			p.SetState(455)
			p.SqlConstraintDefinition()
		}

	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserSQL_WITHOUT_ROWID {
		{
			p.SetState(458)
			p.SqlWithoutRowId()
		}

	}
	{
		p.SetState(461)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(462)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISqlTableFieldDefinitionContext is an interface to support dynamic dispatch.
type ISqlTableFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlTableFieldDefinitionContext differentiates from other interfaces.
	IsSqlTableFieldDefinitionContext()
}

type SqlTableFieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlTableFieldDefinitionContext() *SqlTableFieldDefinitionContext {
	var p = new(SqlTableFieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlTableFieldDefinition
	return p
}

func (*SqlTableFieldDefinitionContext) IsSqlTableFieldDefinitionContext() {}

func NewSqlTableFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlTableFieldDefinitionContext {
	var p = new(SqlTableFieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlTableFieldDefinition

	return p
}

func (s *SqlTableFieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlTableFieldDefinitionContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *SqlTableFieldDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SqlTableFieldDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlTableFieldDefinitionContext) SQL_VIRTUAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserSQL_VIRTUAL, 0)
}

func (s *SqlTableFieldDefinitionContext) SqlConstraint() ISqlConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlConstraintContext)
}

func (s *SqlTableFieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlTableFieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlTableFieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlTableFieldDefinition(s)
	}
}

func (s *SqlTableFieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlTableFieldDefinition(s)
	}
}

func (s *SqlTableFieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlTableFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlTableFieldDefinition() (localctx ISqlTableFieldDefinitionContext) {
	localctx = NewSqlTableFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ZserioParserRULE_sqlTableFieldDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserSQL_VIRTUAL {
		{
			p.SetState(464)
			p.Match(ZserioParserSQL_VIRTUAL)
		}

	}
	{
		p.SetState(467)
		p.TypeInstantiation()
	}
	{
		p.SetState(468)
		p.Id()
	}
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserSQL {
		{
			p.SetState(469)
			p.SqlConstraint()
		}

	}
	{
		p.SetState(472)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISqlConstraintDefinitionContext is an interface to support dynamic dispatch.
type ISqlConstraintDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlConstraintDefinitionContext differentiates from other interfaces.
	IsSqlConstraintDefinitionContext()
}

type SqlConstraintDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlConstraintDefinitionContext() *SqlConstraintDefinitionContext {
	var p = new(SqlConstraintDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlConstraintDefinition
	return p
}

func (*SqlConstraintDefinitionContext) IsSqlConstraintDefinitionContext() {}

func NewSqlConstraintDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlConstraintDefinitionContext {
	var p = new(SqlConstraintDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlConstraintDefinition

	return p
}

func (s *SqlConstraintDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlConstraintDefinitionContext) SqlConstraint() ISqlConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISqlConstraintContext)
}

func (s *SqlConstraintDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlConstraintDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlConstraintDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlConstraintDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlConstraintDefinition(s)
	}
}

func (s *SqlConstraintDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlConstraintDefinition(s)
	}
}

func (s *SqlConstraintDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlConstraintDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlConstraintDefinition() (localctx ISqlConstraintDefinitionContext) {
	localctx = NewSqlConstraintDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ZserioParserRULE_sqlConstraintDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.SqlConstraint()
	}
	{
		p.SetState(475)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISqlConstraintContext is an interface to support dynamic dispatch.
type ISqlConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlConstraintContext differentiates from other interfaces.
	IsSqlConstraintContext()
}

type SqlConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlConstraintContext() *SqlConstraintContext {
	var p = new(SqlConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlConstraint
	return p
}

func (*SqlConstraintContext) IsSqlConstraintContext() {}

func NewSqlConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlConstraintContext {
	var p = new(SqlConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlConstraint

	return p
}

func (s *SqlConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlConstraintContext) SQL() antlr.TerminalNode {
	return s.GetToken(ZserioParserSQL, 0)
}

func (s *SqlConstraintContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqlConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlConstraint(s)
	}
}

func (s *SqlConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlConstraint(s)
	}
}

func (s *SqlConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlConstraint() (localctx ISqlConstraintContext) {
	localctx = NewSqlConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ZserioParserRULE_sqlConstraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Match(ZserioParserSQL)
	}
	{
		p.SetState(478)
		p.expression(0)
	}



	return localctx
}


// ISqlWithoutRowIdContext is an interface to support dynamic dispatch.
type ISqlWithoutRowIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlWithoutRowIdContext differentiates from other interfaces.
	IsSqlWithoutRowIdContext()
}

type SqlWithoutRowIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlWithoutRowIdContext() *SqlWithoutRowIdContext {
	var p = new(SqlWithoutRowIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlWithoutRowId
	return p
}

func (*SqlWithoutRowIdContext) IsSqlWithoutRowIdContext() {}

func NewSqlWithoutRowIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlWithoutRowIdContext {
	var p = new(SqlWithoutRowIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlWithoutRowId

	return p
}

func (s *SqlWithoutRowIdContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlWithoutRowIdContext) SQL_WITHOUT_ROWID() antlr.TerminalNode {
	return s.GetToken(ZserioParserSQL_WITHOUT_ROWID, 0)
}

func (s *SqlWithoutRowIdContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlWithoutRowIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlWithoutRowIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlWithoutRowIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlWithoutRowId(s)
	}
}

func (s *SqlWithoutRowIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlWithoutRowId(s)
	}
}

func (s *SqlWithoutRowIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlWithoutRowId(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlWithoutRowId() (localctx ISqlWithoutRowIdContext) {
	localctx = NewSqlWithoutRowIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ZserioParserRULE_sqlWithoutRowId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(ZserioParserSQL_WITHOUT_ROWID)
	}
	{
		p.SetState(481)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISqlDatabaseDefinitionContext is an interface to support dynamic dispatch.
type ISqlDatabaseDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlDatabaseDefinitionContext differentiates from other interfaces.
	IsSqlDatabaseDefinitionContext()
}

type SqlDatabaseDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlDatabaseDefinitionContext() *SqlDatabaseDefinitionContext {
	var p = new(SqlDatabaseDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlDatabaseDefinition
	return p
}

func (*SqlDatabaseDefinitionContext) IsSqlDatabaseDefinitionContext() {}

func NewSqlDatabaseDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlDatabaseDefinitionContext {
	var p = new(SqlDatabaseDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlDatabaseDefinition

	return p
}

func (s *SqlDatabaseDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlDatabaseDefinitionContext) SQL_DATABASE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSQL_DATABASE, 0)
}

func (s *SqlDatabaseDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SqlDatabaseDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *SqlDatabaseDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *SqlDatabaseDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlDatabaseDefinitionContext) AllSqlDatabaseFieldDefinition() []ISqlDatabaseFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISqlDatabaseFieldDefinitionContext)(nil)).Elem())
	var tst = make([]ISqlDatabaseFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISqlDatabaseFieldDefinitionContext)
		}
	}

	return tst
}

func (s *SqlDatabaseDefinitionContext) SqlDatabaseFieldDefinition(i int) ISqlDatabaseFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISqlDatabaseFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISqlDatabaseFieldDefinitionContext)
}

func (s *SqlDatabaseDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlDatabaseDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlDatabaseDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlDatabaseDefinition(s)
	}
}

func (s *SqlDatabaseDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlDatabaseDefinition(s)
	}
}

func (s *SqlDatabaseDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlDatabaseDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlDatabaseDefinition() (localctx ISqlDatabaseDefinitionContext) {
	localctx = NewSqlDatabaseDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ZserioParserRULE_sqlDatabaseDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(ZserioParserSQL_DATABASE)
	}
	{
		p.SetState(484)
		p.Id()
	}
	{
		p.SetState(485)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(486)
			p.SqlDatabaseFieldDefinition()
		}


		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(491)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(492)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ISqlDatabaseFieldDefinitionContext is an interface to support dynamic dispatch.
type ISqlDatabaseFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlDatabaseFieldDefinitionContext differentiates from other interfaces.
	IsSqlDatabaseFieldDefinitionContext()
}

type SqlDatabaseFieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlDatabaseFieldDefinitionContext() *SqlDatabaseFieldDefinitionContext {
	var p = new(SqlDatabaseFieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_sqlDatabaseFieldDefinition
	return p
}

func (*SqlDatabaseFieldDefinitionContext) IsSqlDatabaseFieldDefinitionContext() {}

func NewSqlDatabaseFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlDatabaseFieldDefinitionContext {
	var p = new(SqlDatabaseFieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_sqlDatabaseFieldDefinition

	return p
}

func (s *SqlDatabaseFieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlDatabaseFieldDefinitionContext) TypeInstantiation() ITypeInstantiationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeInstantiationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeInstantiationContext)
}

func (s *SqlDatabaseFieldDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SqlDatabaseFieldDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *SqlDatabaseFieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlDatabaseFieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SqlDatabaseFieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterSqlDatabaseFieldDefinition(s)
	}
}

func (s *SqlDatabaseFieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitSqlDatabaseFieldDefinition(s)
	}
}

func (s *SqlDatabaseFieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitSqlDatabaseFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) SqlDatabaseFieldDefinition() (localctx ISqlDatabaseFieldDefinitionContext) {
	localctx = NewSqlDatabaseFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ZserioParserRULE_sqlDatabaseFieldDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.TypeInstantiation()
	}
	{
		p.SetState(495)
		p.Id()
	}
	{
		p.SetState(496)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IServiceDefinitionContext is an interface to support dynamic dispatch.
type IServiceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceDefinitionContext differentiates from other interfaces.
	IsServiceDefinitionContext()
}

type ServiceDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefinitionContext() *ServiceDefinitionContext {
	var p = new(ServiceDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_serviceDefinition
	return p
}

func (*ServiceDefinitionContext) IsServiceDefinitionContext() {}

func NewServiceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefinitionContext {
	var p = new(ServiceDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_serviceDefinition

	return p
}

func (s *ServiceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefinitionContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSERVICE, 0)
}

func (s *ServiceDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ServiceDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *ServiceDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *ServiceDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ServiceDefinitionContext) AllServiceMethodDefinition() []IServiceMethodDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceMethodDefinitionContext)(nil)).Elem())
	var tst = make([]IServiceMethodDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceMethodDefinitionContext)
		}
	}

	return tst
}

func (s *ServiceDefinitionContext) ServiceMethodDefinition(i int) IServiceMethodDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceMethodDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceMethodDefinitionContext)
}

func (s *ServiceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ServiceDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterServiceDefinition(s)
	}
}

func (s *ServiceDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitServiceDefinition(s)
	}
}

func (s *ServiceDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitServiceDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ServiceDefinition() (localctx IServiceDefinitionContext) {
	localctx = NewServiceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ZserioParserRULE_serviceDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.Match(ZserioParserSERVICE)
	}
	{
		p.SetState(499)
		p.Id()
	}
	{
		p.SetState(500)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (ZserioParserBIT_FIELD - 33)) | (1 << (ZserioParserBOOL - 33)) | (1 << (ZserioParserEXTERN - 33)) | (1 << (ZserioParserFLOAT16 - 33)) | (1 << (ZserioParserFLOAT32 - 33)) | (1 << (ZserioParserFLOAT64 - 33)) | (1 << (ZserioParserINT_FIELD - 33)) | (1 << (ZserioParserINT16 - 33)) | (1 << (ZserioParserINT32 - 33)) | (1 << (ZserioParserINT64 - 33)) | (1 << (ZserioParserINT8 - 33)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (ZserioParserSTRING - 74)) | (1 << (ZserioParserUINT16 - 74)) | (1 << (ZserioParserUINT32 - 74)) | (1 << (ZserioParserUINT64 - 74)) | (1 << (ZserioParserUINT8 - 74)) | (1 << (ZserioParserVARINT - 74)) | (1 << (ZserioParserVARINT16 - 74)) | (1 << (ZserioParserVARINT32 - 74)) | (1 << (ZserioParserVARINT64 - 74)) | (1 << (ZserioParserVARSIZE - 74)) | (1 << (ZserioParserVARUINT - 74)) | (1 << (ZserioParserVARUINT16 - 74)) | (1 << (ZserioParserVARUINT32 - 74)) | (1 << (ZserioParserVARUINT64 - 74)))) != 0) || _la == ZserioParserID {
		{
			p.SetState(501)
			p.ServiceMethodDefinition()
		}


		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(507)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(508)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IServiceMethodDefinitionContext is an interface to support dynamic dispatch.
type IServiceMethodDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceMethodDefinitionContext differentiates from other interfaces.
	IsServiceMethodDefinitionContext()
}

type ServiceMethodDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceMethodDefinitionContext() *ServiceMethodDefinitionContext {
	var p = new(ServiceMethodDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_serviceMethodDefinition
	return p
}

func (*ServiceMethodDefinitionContext) IsServiceMethodDefinitionContext() {}

func NewServiceMethodDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceMethodDefinitionContext {
	var p = new(ServiceMethodDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_serviceMethodDefinition

	return p
}

func (s *ServiceMethodDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceMethodDefinitionContext) AllTypeReference() []ITypeReferenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem())
	var tst = make([]ITypeReferenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeReferenceContext)
		}
	}

	return tst
}

func (s *ServiceMethodDefinitionContext) TypeReference(i int) ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ServiceMethodDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ServiceMethodDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *ServiceMethodDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *ServiceMethodDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *ServiceMethodDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceMethodDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ServiceMethodDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterServiceMethodDefinition(s)
	}
}

func (s *ServiceMethodDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitServiceMethodDefinition(s)
	}
}

func (s *ServiceMethodDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitServiceMethodDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ServiceMethodDefinition() (localctx IServiceMethodDefinitionContext) {
	localctx = NewServiceMethodDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ZserioParserRULE_serviceMethodDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.TypeReference()
	}
	{
		p.SetState(511)
		p.Id()
	}
	{
		p.SetState(512)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(513)
		p.TypeReference()
	}
	{
		p.SetState(514)
		p.Match(ZserioParserRPAREN)
	}
	{
		p.SetState(515)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IPubsubDefinitionContext is an interface to support dynamic dispatch.
type IPubsubDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPubsubDefinitionContext differentiates from other interfaces.
	IsPubsubDefinitionContext()
}

type PubsubDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPubsubDefinitionContext() *PubsubDefinitionContext {
	var p = new(PubsubDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_pubsubDefinition
	return p
}

func (*PubsubDefinitionContext) IsPubsubDefinitionContext() {}

func NewPubsubDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PubsubDefinitionContext {
	var p = new(PubsubDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_pubsubDefinition

	return p
}

func (s *PubsubDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PubsubDefinitionContext) PUBSUB() antlr.TerminalNode {
	return s.GetToken(ZserioParserPUBSUB, 0)
}

func (s *PubsubDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PubsubDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *PubsubDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *PubsubDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *PubsubDefinitionContext) AllPubsubMessageDefinition() []IPubsubMessageDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPubsubMessageDefinitionContext)(nil)).Elem())
	var tst = make([]IPubsubMessageDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPubsubMessageDefinitionContext)
		}
	}

	return tst
}

func (s *PubsubDefinitionContext) PubsubMessageDefinition(i int) IPubsubMessageDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPubsubMessageDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPubsubMessageDefinitionContext)
}

func (s *PubsubDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PubsubDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PubsubDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterPubsubDefinition(s)
	}
}

func (s *PubsubDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitPubsubDefinition(s)
	}
}

func (s *PubsubDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitPubsubDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) PubsubDefinition() (localctx IPubsubDefinitionContext) {
	localctx = NewPubsubDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ZserioParserRULE_pubsubDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(ZserioParserPUBSUB)
	}
	{
		p.SetState(518)
		p.Id()
	}
	{
		p.SetState(519)
		p.Match(ZserioParserLBRACE)
	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 64)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 64))) & ((1 << (ZserioParserPUBLISH - 64)) | (1 << (ZserioParserSUBSCRIBE - 64)) | (1 << (ZserioParserTOPIC - 64)))) != 0) {
		{
			p.SetState(520)
			p.PubsubMessageDefinition()
		}


		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(526)
		p.Match(ZserioParserRBRACE)
	}
	{
		p.SetState(527)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IPubsubMessageDefinitionContext is an interface to support dynamic dispatch.
type IPubsubMessageDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPubsubMessageDefinitionContext differentiates from other interfaces.
	IsPubsubMessageDefinitionContext()
}

type PubsubMessageDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPubsubMessageDefinitionContext() *PubsubMessageDefinitionContext {
	var p = new(PubsubMessageDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_pubsubMessageDefinition
	return p
}

func (*PubsubMessageDefinitionContext) IsPubsubMessageDefinitionContext() {}

func NewPubsubMessageDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PubsubMessageDefinitionContext {
	var p = new(PubsubMessageDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_pubsubMessageDefinition

	return p
}

func (s *PubsubMessageDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PubsubMessageDefinitionContext) TopicDefinition() ITopicDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopicDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITopicDefinitionContext)
}

func (s *PubsubMessageDefinitionContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *PubsubMessageDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PubsubMessageDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *PubsubMessageDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PubsubMessageDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PubsubMessageDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterPubsubMessageDefinition(s)
	}
}

func (s *PubsubMessageDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitPubsubMessageDefinition(s)
	}
}

func (s *PubsubMessageDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitPubsubMessageDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) PubsubMessageDefinition() (localctx IPubsubMessageDefinitionContext) {
	localctx = NewPubsubMessageDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ZserioParserRULE_pubsubMessageDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.TopicDefinition()
	}
	{
		p.SetState(530)
		p.TypeReference()
	}
	{
		p.SetState(531)
		p.Id()
	}
	{
		p.SetState(532)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// ITopicDefinitionContext is an interface to support dynamic dispatch.
type ITopicDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopicDefinitionContext differentiates from other interfaces.
	IsTopicDefinitionContext()
}

type TopicDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopicDefinitionContext() *TopicDefinitionContext {
	var p = new(TopicDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_topicDefinition
	return p
}

func (*TopicDefinitionContext) IsTopicDefinitionContext() {}

func NewTopicDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopicDefinitionContext {
	var p = new(TopicDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_topicDefinition

	return p
}

func (s *TopicDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TopicDefinitionContext) TOPIC() antlr.TerminalNode {
	return s.GetToken(ZserioParserTOPIC, 0)
}

func (s *TopicDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *TopicDefinitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TopicDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *TopicDefinitionContext) PUBLISH() antlr.TerminalNode {
	return s.GetToken(ZserioParserPUBLISH, 0)
}

func (s *TopicDefinitionContext) SUBSCRIBE() antlr.TerminalNode {
	return s.GetToken(ZserioParserSUBSCRIBE, 0)
}

func (s *TopicDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopicDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TopicDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTopicDefinition(s)
	}
}

func (s *TopicDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTopicDefinition(s)
	}
}

func (s *TopicDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTopicDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TopicDefinition() (localctx ITopicDefinitionContext) {
	localctx = NewTopicDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ZserioParserRULE_topicDefinition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ZserioParserPUBLISH || _la == ZserioParserSUBSCRIBE {
		{
			p.SetState(534)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ZserioParserPUBLISH || _la == ZserioParserSUBSCRIBE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(537)
		p.Match(ZserioParserTOPIC)
	}
	{
		p.SetState(538)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(539)
		p.expression(0)
	}
	{
		p.SetState(540)
		p.Match(ZserioParserRPAREN)
	}



	return localctx
}


// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ZserioParserFUNCTION, 0)
}

func (s *FunctionDefinitionContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *FunctionDefinitionContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *FunctionDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *FunctionDefinitionContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ZserioParserRULE_functionDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		p.Match(ZserioParserFUNCTION)
	}
	{
		p.SetState(543)
		p.FunctionType()
	}
	{
		p.SetState(544)
		p.FunctionName()
	}
	{
		p.SetState(545)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(546)
		p.Match(ZserioParserRPAREN)
	}
	{
		p.SetState(547)
		p.FunctionBody()
	}



	return localctx
}


// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ZserioParserRULE_functionType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.TypeReference()
	}



	return localctx
}


// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (s *FunctionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFunctionName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ZserioParserRULE_functionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Id()
	}



	return localctx
}


// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACE, 0)
}

func (s *FunctionBodyContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRETURN, 0)
}

func (s *FunctionBodyContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionBodyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *FunctionBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACE, 0)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ZserioParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		p.Match(ZserioParserLBRACE)
	}
	{
		p.SetState(554)
		p.Match(ZserioParserRETURN)
	}
	{
		p.SetState(555)
		p.expression(0)
	}
	{
		p.SetState(556)
		p.Match(ZserioParserSEMICOLON)
	}
	{
		p.SetState(557)
		p.Match(ZserioParserRBRACE)
	}



	return localctx
}


// ITypeParametersContext is an interface to support dynamic dispatch.
type ITypeParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParametersContext differentiates from other interfaces.
	IsTypeParametersContext()
}

type TypeParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParametersContext() *TypeParametersContext {
	var p = new(TypeParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeParameters
	return p
}

func (*TypeParametersContext) IsTypeParametersContext() {}

func NewTypeParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParametersContext {
	var p = new(TypeParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeParameters

	return p
}

func (s *TypeParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *TypeParametersContext) AllParameterDefinition() []IParameterDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterDefinitionContext)(nil)).Elem())
	var tst = make([]IParameterDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterDefinitionContext)
		}
	}

	return tst
}

func (s *TypeParametersContext) ParameterDefinition(i int) IParameterDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterDefinitionContext)
}

func (s *TypeParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *TypeParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *TypeParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *TypeParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeParameters(s)
	}
}

func (s *TypeParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeParameters(s)
	}
}

func (s *TypeParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeParameters(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeParameters() (localctx ITypeParametersContext) {
	localctx = NewTypeParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ZserioParserRULE_typeParameters)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(560)
		p.ParameterDefinition()
	}
	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserCOMMA {
		{
			p.SetState(561)
			p.Match(ZserioParserCOMMA)
		}
		{
			p.SetState(562)
			p.ParameterDefinition()
		}


		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(568)
		p.Match(ZserioParserRPAREN)
	}



	return localctx
}


// IParameterDefinitionContext is an interface to support dynamic dispatch.
type IParameterDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterDefinitionContext differentiates from other interfaces.
	IsParameterDefinitionContext()
}

type ParameterDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDefinitionContext() *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_parameterDefinition
	return p
}

func (*ParameterDefinitionContext) IsParameterDefinitionContext() {}

func NewParameterDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_parameterDefinition

	return p
}

func (s *ParameterDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDefinitionContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ParameterDefinitionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ParameterDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParameterDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterParameterDefinition(s)
	}
}

func (s *ParameterDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitParameterDefinition(s)
	}
}

func (s *ParameterDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitParameterDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ParameterDefinition() (localctx IParameterDefinitionContext) {
	localctx = NewParameterDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ZserioParserRULE_parameterDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(570)
		p.TypeReference()
	}
	{
		p.SetState(571)
		p.Id()
	}



	return localctx
}


// ITemplateParametersContext is an interface to support dynamic dispatch.
type ITemplateParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateParametersContext differentiates from other interfaces.
	IsTemplateParametersContext()
}

type TemplateParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateParametersContext() *TemplateParametersContext {
	var p = new(TemplateParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_templateParameters
	return p
}

func (*TemplateParametersContext) IsTemplateParametersContext() {}

func NewTemplateParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateParametersContext {
	var p = new(TemplateParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_templateParameters

	return p
}

func (s *TemplateParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateParametersContext) LT() antlr.TerminalNode {
	return s.GetToken(ZserioParserLT, 0)
}

func (s *TemplateParametersContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *TemplateParametersContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *TemplateParametersContext) GT() antlr.TerminalNode {
	return s.GetToken(ZserioParserGT, 0)
}

func (s *TemplateParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *TemplateParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *TemplateParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemplateParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTemplateParameters(s)
	}
}

func (s *TemplateParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTemplateParameters(s)
	}
}

func (s *TemplateParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTemplateParameters(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TemplateParameters() (localctx ITemplateParametersContext) {
	localctx = NewTemplateParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ZserioParserRULE_templateParameters)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(ZserioParserLT)
	}
	{
		p.SetState(574)
		p.Id()
	}
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserCOMMA {
		{
			p.SetState(575)
			p.Match(ZserioParserCOMMA)
		}
		{
			p.SetState(576)
			p.Id()
		}


		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(582)
		p.Match(ZserioParserGT)
	}



	return localctx
}


// ITemplateArgumentsContext is an interface to support dynamic dispatch.
type ITemplateArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateArgumentsContext differentiates from other interfaces.
	IsTemplateArgumentsContext()
}

type TemplateArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateArgumentsContext() *TemplateArgumentsContext {
	var p = new(TemplateArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_templateArguments
	return p
}

func (*TemplateArgumentsContext) IsTemplateArgumentsContext() {}

func NewTemplateArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateArgumentsContext {
	var p = new(TemplateArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_templateArguments

	return p
}

func (s *TemplateArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateArgumentsContext) LT() antlr.TerminalNode {
	return s.GetToken(ZserioParserLT, 0)
}

func (s *TemplateArgumentsContext) AllTemplateArgument() []ITemplateArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITemplateArgumentContext)(nil)).Elem())
	var tst = make([]ITemplateArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITemplateArgumentContext)
		}
	}

	return tst
}

func (s *TemplateArgumentsContext) TemplateArgument(i int) ITemplateArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITemplateArgumentContext)
}

func (s *TemplateArgumentsContext) GT() antlr.TerminalNode {
	return s.GetToken(ZserioParserGT, 0)
}

func (s *TemplateArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *TemplateArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *TemplateArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemplateArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTemplateArguments(s)
	}
}

func (s *TemplateArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTemplateArguments(s)
	}
}

func (s *TemplateArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTemplateArguments(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TemplateArguments() (localctx ITemplateArgumentsContext) {
	localctx = NewTemplateArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ZserioParserRULE_templateArguments)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.Match(ZserioParserLT)
	}
	{
		p.SetState(585)
		p.TemplateArgument()
	}
	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserCOMMA {
		{
			p.SetState(586)
			p.Match(ZserioParserCOMMA)
		}
		{
			p.SetState(587)
			p.TemplateArgument()
		}


		p.SetState(592)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(593)
		p.Match(ZserioParserGT)
	}



	return localctx
}


// ITemplateArgumentContext is an interface to support dynamic dispatch.
type ITemplateArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateArgumentContext differentiates from other interfaces.
	IsTemplateArgumentContext()
}

type TemplateArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateArgumentContext() *TemplateArgumentContext {
	var p = new(TemplateArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_templateArgument
	return p
}

func (*TemplateArgumentContext) IsTemplateArgumentContext() {}

func NewTemplateArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateArgumentContext {
	var p = new(TemplateArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_templateArgument

	return p
}

func (s *TemplateArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateArgumentContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *TemplateArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemplateArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTemplateArgument(s)
	}
}

func (s *TemplateArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTemplateArgument(s)
	}
}

func (s *TemplateArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTemplateArgument(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TemplateArgument() (localctx ITemplateArgumentContext) {
	localctx = NewTemplateArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ZserioParserRULE_templateArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(595)
		p.TypeReference()
	}



	return localctx
}


// IInstantiateDeclarationContext is an interface to support dynamic dispatch.
type IInstantiateDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstantiateDeclarationContext differentiates from other interfaces.
	IsInstantiateDeclarationContext()
}

type InstantiateDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstantiateDeclarationContext() *InstantiateDeclarationContext {
	var p = new(InstantiateDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_instantiateDeclaration
	return p
}

func (*InstantiateDeclarationContext) IsInstantiateDeclarationContext() {}

func NewInstantiateDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstantiateDeclarationContext {
	var p = new(InstantiateDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_instantiateDeclaration

	return p
}

func (s *InstantiateDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InstantiateDeclarationContext) INSTANTIATE() antlr.TerminalNode {
	return s.GetToken(ZserioParserINSTANTIATE, 0)
}

func (s *InstantiateDeclarationContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *InstantiateDeclarationContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *InstantiateDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserSEMICOLON, 0)
}

func (s *InstantiateDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstantiateDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InstantiateDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterInstantiateDeclaration(s)
	}
}

func (s *InstantiateDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitInstantiateDeclaration(s)
	}
}

func (s *InstantiateDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitInstantiateDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) InstantiateDeclaration() (localctx IInstantiateDeclarationContext) {
	localctx = NewInstantiateDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ZserioParserRULE_instantiateDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		p.Match(ZserioParserINSTANTIATE)
	}
	{
		p.SetState(598)
		p.TypeReference()
	}
	{
		p.SetState(599)
		p.Id()
	}
	{
		p.SetState(600)
		p.Match(ZserioParserSEMICOLON)
	}



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type BitwiseXorExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewBitwiseXorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseXorExpressionContext {
	var p = new(BitwiseXorExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *BitwiseXorExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *BitwiseXorExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BitwiseXorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseXorExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitwiseXorExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitwiseXorExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(ZserioParserXOR, 0)
}


func (s *BitwiseXorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBitwiseXorExpression(s)
	}
}

func (s *BitwiseXorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBitwiseXorExpression(s)
	}
}

func (s *BitwiseXorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBitwiseXorExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type DotExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewDotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotExpressionContext {
	var p = new(DotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *DotExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *DotExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *DotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DotExpressionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *DotExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ZserioParserDOT, 0)
}


func (s *DotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterDotExpression(s)
	}
}

func (s *DotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitDotExpression(s)
	}
}

func (s *DotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitDotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ValueofExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewValueofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueofExpressionContext {
	var p = new(ValueofExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *ValueofExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *ValueofExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ValueofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueofExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *ValueofExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ValueofExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *ValueofExpressionContext) VALUEOF() antlr.TerminalNode {
	return s.GetToken(ZserioParserVALUEOF, 0)
}


func (s *ValueofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterValueofExpression(s)
	}
}

func (s *ValueofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitValueofExpression(s)
	}
}

func (s *ValueofExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitValueofExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ShiftExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *ShiftExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *ShiftExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ShiftExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShiftExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserGT)
}

func (s *ShiftExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserGT, i)
}

func (s *ShiftExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ZserioParserLSHIFT, 0)
}


func (s *ShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitShiftExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *ArrayExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *ArrayExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZserioParserRBRACKET, 0)
}

func (s *ArrayExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZserioParserLBRACKET, 0)
}


func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitArrayExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NumbitsExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewNumbitsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumbitsExpressionContext {
	var p = new(NumbitsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *NumbitsExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *NumbitsExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *NumbitsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumbitsExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *NumbitsExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NumbitsExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *NumbitsExpressionContext) NUMBITS() antlr.TerminalNode {
	return s.GetToken(ZserioParserNUMBITS, 0)
}


func (s *NumbitsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterNumbitsExpression(s)
	}
}

func (s *NumbitsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitNumbitsExpression(s)
	}
}

func (s *NumbitsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitNumbitsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type AdditiveExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *AdditiveExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *AdditiveExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditiveExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZserioParserPLUS, 0)
}

func (s *AdditiveExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZserioParserMINUS, 0)
}


func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type RelationalExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *RelationalExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *RelationalExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelationalExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(ZserioParserLT, 0)
}

func (s *RelationalExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(ZserioParserLE, 0)
}

func (s *RelationalExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(ZserioParserGT, 0)
}

func (s *RelationalExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(ZserioParserGE, 0)
}


func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LengthofExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewLengthofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthofExpressionContext {
	var p = new(LengthofExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *LengthofExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *LengthofExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *LengthofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthofExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *LengthofExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthofExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *LengthofExpressionContext) LENGTHOF() antlr.TerminalNode {
	return s.GetToken(ZserioParserLENGTHOF, 0)
}


func (s *LengthofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLengthofExpression(s)
	}
}

func (s *LengthofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLengthofExpression(s)
	}
}

func (s *LengthofExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLengthofExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}


func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type MultiplicativeExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *MultiplicativeExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *MultiplicativeExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicativeExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(ZserioParserMULTIPLY, 0)
}

func (s *MultiplicativeExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ZserioParserDIVIDE, 0)
}

func (s *MultiplicativeExpressionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(ZserioParserMODULO, 0)
}


func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LogicalOrExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *LogicalOrExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *LogicalOrExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalOrExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(ZserioParserLOGICAL_OR, 0)
}


func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type BitwiseOrExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewBitwiseOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseOrExpressionContext {
	var p = new(BitwiseOrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *BitwiseOrExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *BitwiseOrExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BitwiseOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseOrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitwiseOrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitwiseOrExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(ZserioParserOR, 0)
}


func (s *BitwiseOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBitwiseOrExpression(s)
	}
}

func (s *BitwiseOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBitwiseOrExpression(s)
	}
}

func (s *BitwiseOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBitwiseOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParenthesizedExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *ParenthesizedExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *ParenthesizedExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *ParenthesizedExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}


func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitParenthesizedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type BitwiseAndExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewBitwiseAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseAndExpressionContext {
	var p = new(BitwiseAndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *BitwiseAndExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *BitwiseAndExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BitwiseAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseAndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BitwiseAndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitwiseAndExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(ZserioParserAND, 0)
}


func (s *BitwiseAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBitwiseAndExpression(s)
	}
}

func (s *BitwiseAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBitwiseAndExpression(s)
	}
}

func (s *BitwiseAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBitwiseAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type EqualityExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *EqualityExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *EqualityExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZserioParserEQ, 0)
}

func (s *EqualityExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(ZserioParserNE, 0)
}


func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LogicalAndExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *LogicalAndExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *LogicalAndExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalAndExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(ZserioParserLOGICAL_AND, 0)
}


func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type FunctionCallExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *FunctionCallExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *FunctionCallExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *FunctionCallExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}


func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type IndexExpressionContext struct {
	*ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) INDEX() antlr.TerminalNode {
	return s.GetToken(ZserioParserINDEX, 0)
}


func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (s *IndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type UnaryExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *UnaryExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *UnaryExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZserioParserPLUS, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZserioParserMINUS, 0)
}

func (s *UnaryExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(ZserioParserBANG, 0)
}

func (s *UnaryExpressionContext) TILDE() antlr.TerminalNode {
	return s.GetToken(ZserioParserTILDE, 0)
}


func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LiteralExpressionContext struct {
	*ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}


func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type TernaryExpressionContext struct {
	*ExpressionContext
	operator antlr.Token
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *TernaryExpressionContext) GetOperator() antlr.Token { return s.operator }


func (s *TernaryExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *TernaryExpressionContext) QUESTIONMARK() antlr.TerminalNode {
	return s.GetToken(ZserioParserQUESTIONMARK, 0)
}


func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTernaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *ZserioParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ZserioParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 104
	p.EnterRecursionRule(localctx, 104, ZserioParserRULE_expression, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(627)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserLPAREN:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(603)

			var _m = p.Match(ZserioParserLPAREN)

			localctx.(*ParenthesizedExpressionContext).operator = _m
		}
		{
			p.SetState(604)
			p.expression(0)
		}
		{
			p.SetState(605)
			p.Match(ZserioParserRPAREN)
		}


	case ZserioParserLENGTHOF:
		localctx = NewLengthofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(607)

			var _m = p.Match(ZserioParserLENGTHOF)

			localctx.(*LengthofExpressionContext).operator = _m
		}
		{
			p.SetState(608)
			p.Match(ZserioParserLPAREN)
		}
		{
			p.SetState(609)
			p.expression(0)
		}
		{
			p.SetState(610)
			p.Match(ZserioParserRPAREN)
		}


	case ZserioParserVALUEOF:
		localctx = NewValueofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(612)

			var _m = p.Match(ZserioParserVALUEOF)

			localctx.(*ValueofExpressionContext).operator = _m
		}
		{
			p.SetState(613)
			p.Match(ZserioParserLPAREN)
		}
		{
			p.SetState(614)
			p.expression(0)
		}
		{
			p.SetState(615)
			p.Match(ZserioParserRPAREN)
		}


	case ZserioParserNUMBITS:
		localctx = NewNumbitsExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(617)

			var _m = p.Match(ZserioParserNUMBITS)

			localctx.(*NumbitsExpressionContext).operator = _m
		}
		{
			p.SetState(618)
			p.Match(ZserioParserLPAREN)
		}
		{
			p.SetState(619)
			p.expression(0)
		}
		{
			p.SetState(620)
			p.Match(ZserioParserRPAREN)
		}


	case ZserioParserBANG, ZserioParserMINUS, ZserioParserPLUS, ZserioParserTILDE:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(622)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZserioParserBANG) | (1 << ZserioParserMINUS) | (1 << ZserioParserPLUS) | (1 << ZserioParserTILDE))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(623)
			p.expression(15)
		}


	case ZserioParserBOOL_LITERAL, ZserioParserSTRING_LITERAL, ZserioParserBINARY_LITERAL, ZserioParserOCTAL_LITERAL, ZserioParserHEXADECIMAL_LITERAL, ZserioParserDOUBLE_LITERAL, ZserioParserFLOAT_LITERAL, ZserioParserDECIMAL_LITERAL:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(624)
			p.Literal()
		}


	case ZserioParserINDEX:
		localctx = NewIndexExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(625)
			p.Match(ZserioParserINDEX)
		}


	case ZserioParserID:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(626)
			p.Id()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(680)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(629)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(630)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicativeExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZserioParserDIVIDE) | (1 << ZserioParserMODULO) | (1 << ZserioParserMULTIPLY))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicativeExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(631)
					p.expression(15)
				}


			case 2:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(632)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(633)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZserioParserMINUS || _la == ZserioParserPLUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(634)
					p.expression(14)
				}


			case 3:
				localctx = NewShiftExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(635)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(639)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case ZserioParserLSHIFT:
					{
						p.SetState(636)

						var _m = p.Match(ZserioParserLSHIFT)

						localctx.(*ShiftExpressionContext).operator = _m
					}


				case ZserioParserGT:
					{
						p.SetState(637)

						var _m = p.Match(ZserioParserGT)

						localctx.(*ShiftExpressionContext).operator = _m
					}
					{
						p.SetState(638)
						p.Match(ZserioParserGT)
					}



				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(641)
					p.expression(13)
				}


			case 4:
				localctx = NewRelationalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(642)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(643)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZserioParserGE) | (1 << ZserioParserGT) | (1 << ZserioParserLE) | (1 << ZserioParserLT))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(644)
					p.expression(12)
				}


			case 5:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(645)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(646)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZserioParserEQ || _la == ZserioParserNE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(647)
					p.expression(11)
				}


			case 6:
				localctx = NewBitwiseAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(648)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(649)

					var _m = p.Match(ZserioParserAND)

					localctx.(*BitwiseAndExpressionContext).operator = _m
				}
				{
					p.SetState(650)
					p.expression(10)
				}


			case 7:
				localctx = NewBitwiseXorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(651)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(652)

					var _m = p.Match(ZserioParserXOR)

					localctx.(*BitwiseXorExpressionContext).operator = _m
				}
				{
					p.SetState(653)
					p.expression(9)
				}


			case 8:
				localctx = NewBitwiseOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(654)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(655)

					var _m = p.Match(ZserioParserOR)

					localctx.(*BitwiseOrExpressionContext).operator = _m
				}
				{
					p.SetState(656)
					p.expression(8)
				}


			case 9:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(657)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(658)

					var _m = p.Match(ZserioParserLOGICAL_AND)

					localctx.(*LogicalAndExpressionContext).operator = _m
				}
				{
					p.SetState(659)
					p.expression(7)
				}


			case 10:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(660)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(661)

					var _m = p.Match(ZserioParserLOGICAL_OR)

					localctx.(*LogicalOrExpressionContext).operator = _m
				}
				{
					p.SetState(662)
					p.expression(6)
				}


			case 11:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(663)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(664)

					var _m = p.Match(ZserioParserQUESTIONMARK)

					localctx.(*TernaryExpressionContext).operator = _m
				}
				{
					p.SetState(665)
					p.expression(0)
				}
				{
					p.SetState(666)
					p.Match(ZserioParserCOLON)
				}
				{
					p.SetState(667)
					p.expression(4)
				}


			case 12:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(669)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(670)
					p.Match(ZserioParserLPAREN)
				}
				{
					p.SetState(671)

					var _m = p.Match(ZserioParserRPAREN)

					localctx.(*FunctionCallExpressionContext).operator = _m
				}


			case 13:
				localctx = NewArrayExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(672)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(673)

					var _m = p.Match(ZserioParserLBRACKET)

					localctx.(*ArrayExpressionContext).operator = _m
				}
				{
					p.SetState(674)
					p.expression(0)
				}
				{
					p.SetState(675)
					p.Match(ZserioParserRBRACKET)
				}


			case 14:
				localctx = NewDotExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZserioParserRULE_expression)
				p.SetState(677)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(678)

					var _m = p.Match(ZserioParserDOT)

					localctx.(*DotExpressionContext).operator = _m
				}
				{
					p.SetState(679)
					p.Id()
				}

			}

		}
		p.SetState(684)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
	}



	return localctx
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserBINARY_LITERAL, 0)
}

func (s *LiteralContext) OCTAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserOCTAL_LITERAL, 0)
}

func (s *LiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserDECIMAL_LITERAL, 0)
}

func (s *LiteralContext) HEXADECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserHEXADECIMAL_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserFLOAT_LITERAL, 0)
}

func (s *LiteralContext) DOUBLE_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserDOUBLE_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ZserioParserRULE_literal)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(685)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 100)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 100))) & ((1 << (ZserioParserBOOL_LITERAL - 100)) | (1 << (ZserioParserSTRING_LITERAL - 100)) | (1 << (ZserioParserBINARY_LITERAL - 100)) | (1 << (ZserioParserOCTAL_LITERAL - 100)) | (1 << (ZserioParserHEXADECIMAL_LITERAL - 100)) | (1 << (ZserioParserDOUBLE_LITERAL - 100)) | (1 << (ZserioParserFLOAT_LITERAL - 100)) | (1 << (ZserioParserDECIMAL_LITERAL - 100)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(ZserioParserID, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ZserioParserRULE_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(687)
		p.Match(ZserioParserID)
	}



	return localctx
}


// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeReference
	return p
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) BuiltinType() IBuiltinTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinTypeContext)
}

func (s *TypeReferenceContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *TypeReferenceContext) TemplateArguments() ITemplateArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateArgumentsContext)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (s *TypeReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeReference(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeReference() (localctx ITypeReferenceContext) {
	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ZserioParserRULE_typeReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(694)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserBIT_FIELD, ZserioParserBOOL, ZserioParserEXTERN, ZserioParserFLOAT16, ZserioParserFLOAT32, ZserioParserFLOAT64, ZserioParserINT_FIELD, ZserioParserINT16, ZserioParserINT32, ZserioParserINT64, ZserioParserINT8, ZserioParserSTRING, ZserioParserUINT16, ZserioParserUINT32, ZserioParserUINT64, ZserioParserUINT8, ZserioParserVARINT, ZserioParserVARINT16, ZserioParserVARINT32, ZserioParserVARINT64, ZserioParserVARSIZE, ZserioParserVARUINT, ZserioParserVARUINT16, ZserioParserVARUINT32, ZserioParserVARUINT64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(689)
			p.BuiltinType()
		}


	case ZserioParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(690)
			p.QualifiedName()
		}
		p.SetState(692)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(691)
				p.TemplateArguments()
			}


		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITypeInstantiationContext is an interface to support dynamic dispatch.
type ITypeInstantiationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeInstantiationContext differentiates from other interfaces.
	IsTypeInstantiationContext()
}

type TypeInstantiationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeInstantiationContext() *TypeInstantiationContext {
	var p = new(TypeInstantiationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeInstantiation
	return p
}

func (*TypeInstantiationContext) IsTypeInstantiationContext() {}

func NewTypeInstantiationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeInstantiationContext {
	var p = new(TypeInstantiationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeInstantiation

	return p
}

func (s *TypeInstantiationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeInstantiationContext) TypeReference() ITypeReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *TypeInstantiationContext) TypeArguments() ITypeArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeArgumentsContext)
}

func (s *TypeInstantiationContext) DynamicLengthArgument() IDynamicLengthArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicLengthArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicLengthArgumentContext)
}

func (s *TypeInstantiationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeInstantiationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeInstantiationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeInstantiation(s)
	}
}

func (s *TypeInstantiationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeInstantiation(s)
	}
}

func (s *TypeInstantiationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeInstantiation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeInstantiation() (localctx ITypeInstantiationContext) {
	localctx = NewTypeInstantiationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ZserioParserRULE_typeInstantiation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(696)
		p.TypeReference()
	}
	p.SetState(699)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserLPAREN:
		{
			p.SetState(697)
			p.TypeArguments()
		}


	case ZserioParserLT:
		{
			p.SetState(698)
			p.DynamicLengthArgument()
		}


	case ZserioParserID:



	default:
	}



	return localctx
}


// IBuiltinTypeContext is an interface to support dynamic dispatch.
type IBuiltinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinTypeContext differentiates from other interfaces.
	IsBuiltinTypeContext()
}

type BuiltinTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinTypeContext() *BuiltinTypeContext {
	var p = new(BuiltinTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_builtinType
	return p
}

func (*BuiltinTypeContext) IsBuiltinTypeContext() {}

func NewBuiltinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinTypeContext {
	var p = new(BuiltinTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_builtinType

	return p
}

func (s *BuiltinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinTypeContext) IntType() IIntTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntTypeContext)
}

func (s *BuiltinTypeContext) VarintType() IVarintTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarintTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarintTypeContext)
}

func (s *BuiltinTypeContext) FixedBitFieldType() IFixedBitFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFixedBitFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFixedBitFieldTypeContext)
}

func (s *BuiltinTypeContext) DynamicBitFieldType() IDynamicBitFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicBitFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicBitFieldTypeContext)
}

func (s *BuiltinTypeContext) BoolType() IBoolTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolTypeContext)
}

func (s *BuiltinTypeContext) StringType() IStringTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *BuiltinTypeContext) FloatType() IFloatTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatTypeContext)
}

func (s *BuiltinTypeContext) ExternType() IExternTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternTypeContext)
}

func (s *BuiltinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BuiltinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBuiltinType(s)
	}
}

func (s *BuiltinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBuiltinType(s)
	}
}

func (s *BuiltinTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBuiltinType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) BuiltinType() (localctx IBuiltinTypeContext) {
	localctx = NewBuiltinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ZserioParserRULE_builtinType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(709)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(701)
			p.IntType()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(702)
			p.VarintType()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(703)
			p.FixedBitFieldType()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(704)
			p.DynamicBitFieldType()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(705)
			p.BoolType()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(706)
			p.StringType()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(707)
			p.FloatType()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(708)
			p.ExternType()
		}

	}


	return localctx
}


// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *QualifiedNameContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserDOT, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ZserioParserRULE_qualifiedName)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(711)
		p.Id()
	}
	p.SetState(716)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserDOT {
		{
			p.SetState(712)
			p.Match(ZserioParserDOT)
		}
		{
			p.SetState(713)
			p.Id()
		}


		p.SetState(718)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ITypeArgumentsContext is an interface to support dynamic dispatch.
type ITypeArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentsContext differentiates from other interfaces.
	IsTypeArgumentsContext()
}

type TypeArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentsContext() *TypeArgumentsContext {
	var p = new(TypeArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeArguments
	return p
}

func (*TypeArgumentsContext) IsTypeArgumentsContext() {}

func NewTypeArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentsContext {
	var p = new(TypeArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeArguments

	return p
}

func (s *TypeArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserLPAREN, 0)
}

func (s *TypeArgumentsContext) AllTypeArgument() []ITypeArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeArgumentContext)(nil)).Elem())
	var tst = make([]ITypeArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeArgumentContext)
		}
	}

	return tst
}

func (s *TypeArgumentsContext) TypeArgument(i int) ITypeArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeArgumentContext)
}

func (s *TypeArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZserioParserRPAREN, 0)
}

func (s *TypeArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZserioParserCOMMA)
}

func (s *TypeArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZserioParserCOMMA, i)
}

func (s *TypeArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeArguments(s)
	}
}

func (s *TypeArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeArguments(s)
	}
}

func (s *TypeArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeArguments(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeArguments() (localctx ITypeArgumentsContext) {
	localctx = NewTypeArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ZserioParserRULE_typeArguments)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(719)
		p.Match(ZserioParserLPAREN)
	}
	{
		p.SetState(720)
		p.TypeArgument()
	}
	p.SetState(725)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZserioParserCOMMA {
		{
			p.SetState(721)
			p.Match(ZserioParserCOMMA)
		}
		{
			p.SetState(722)
			p.TypeArgument()
		}


		p.SetState(727)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(728)
		p.Match(ZserioParserRPAREN)
	}



	return localctx
}


// ITypeArgumentContext is an interface to support dynamic dispatch.
type ITypeArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentContext differentiates from other interfaces.
	IsTypeArgumentContext()
}

type TypeArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentContext() *TypeArgumentContext {
	var p = new(TypeArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_typeArgument
	return p
}

func (*TypeArgumentContext) IsTypeArgumentContext() {}

func NewTypeArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentContext {
	var p = new(TypeArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_typeArgument

	return p
}

func (s *TypeArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentContext) EXPLICIT() antlr.TerminalNode {
	return s.GetToken(ZserioParserEXPLICIT, 0)
}

func (s *TypeArgumentContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *TypeArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterTypeArgument(s)
	}
}

func (s *TypeArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitTypeArgument(s)
	}
}

func (s *TypeArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitTypeArgument(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) TypeArgument() (localctx ITypeArgumentContext) {
	localctx = NewTypeArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ZserioParserRULE_typeArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(733)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZserioParserEXPLICIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(730)
			p.Match(ZserioParserEXPLICIT)
		}
		{
			p.SetState(731)
			p.Id()
		}


	case ZserioParserBANG, ZserioParserLPAREN, ZserioParserMINUS, ZserioParserPLUS, ZserioParserTILDE, ZserioParserINDEX, ZserioParserLENGTHOF, ZserioParserNUMBITS, ZserioParserVALUEOF, ZserioParserBOOL_LITERAL, ZserioParserSTRING_LITERAL, ZserioParserBINARY_LITERAL, ZserioParserOCTAL_LITERAL, ZserioParserHEXADECIMAL_LITERAL, ZserioParserDOUBLE_LITERAL, ZserioParserFLOAT_LITERAL, ZserioParserDECIMAL_LITERAL, ZserioParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(732)
			p.expression(0)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IDynamicLengthArgumentContext is an interface to support dynamic dispatch.
type IDynamicLengthArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDynamicLengthArgumentContext differentiates from other interfaces.
	IsDynamicLengthArgumentContext()
}

type DynamicLengthArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicLengthArgumentContext() *DynamicLengthArgumentContext {
	var p = new(DynamicLengthArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_dynamicLengthArgument
	return p
}

func (*DynamicLengthArgumentContext) IsDynamicLengthArgumentContext() {}

func NewDynamicLengthArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicLengthArgumentContext {
	var p = new(DynamicLengthArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_dynamicLengthArgument

	return p
}

func (s *DynamicLengthArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicLengthArgumentContext) LT() antlr.TerminalNode {
	return s.GetToken(ZserioParserLT, 0)
}

func (s *DynamicLengthArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DynamicLengthArgumentContext) GT() antlr.TerminalNode {
	return s.GetToken(ZserioParserGT, 0)
}

func (s *DynamicLengthArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicLengthArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DynamicLengthArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterDynamicLengthArgument(s)
	}
}

func (s *DynamicLengthArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitDynamicLengthArgument(s)
	}
}

func (s *DynamicLengthArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitDynamicLengthArgument(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) DynamicLengthArgument() (localctx IDynamicLengthArgumentContext) {
	localctx = NewDynamicLengthArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ZserioParserRULE_dynamicLengthArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(735)
		p.Match(ZserioParserLT)
	}
	{
		p.SetState(736)
		p.expression(0)
	}
	{
		p.SetState(737)
		p.Match(ZserioParserGT)
	}



	return localctx
}


// IIntTypeContext is an interface to support dynamic dispatch.
type IIntTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntTypeContext differentiates from other interfaces.
	IsIntTypeContext()
}

type IntTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntTypeContext() *IntTypeContext {
	var p = new(IntTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_intType
	return p
}

func (*IntTypeContext) IsIntTypeContext() {}

func NewIntTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntTypeContext {
	var p = new(IntTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_intType

	return p
}

func (s *IntTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IntTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT8, 0)
}

func (s *IntTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT16, 0)
}

func (s *IntTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT32, 0)
}

func (s *IntTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT64, 0)
}

func (s *IntTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(ZserioParserUINT8, 0)
}

func (s *IntTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(ZserioParserUINT16, 0)
}

func (s *IntTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(ZserioParserUINT32, 0)
}

func (s *IntTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(ZserioParserUINT64, 0)
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) IntType() (localctx IIntTypeContext) {
	localctx = NewIntTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ZserioParserRULE_intType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(739)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 53)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 53))) & ((1 << (ZserioParserINT16 - 53)) | (1 << (ZserioParserINT32 - 53)) | (1 << (ZserioParserINT64 - 53)) | (1 << (ZserioParserINT8 - 53)) | (1 << (ZserioParserUINT16 - 53)) | (1 << (ZserioParserUINT32 - 53)) | (1 << (ZserioParserUINT64 - 53)) | (1 << (ZserioParserUINT8 - 53)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IVarintTypeContext is an interface to support dynamic dispatch.
type IVarintTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarintTypeContext differentiates from other interfaces.
	IsVarintTypeContext()
}

type VarintTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarintTypeContext() *VarintTypeContext {
	var p = new(VarintTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_varintType
	return p
}

func (*VarintTypeContext) IsVarintTypeContext() {}

func NewVarintTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarintTypeContext {
	var p = new(VarintTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_varintType

	return p
}

func (s *VarintTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarintTypeContext) VARINT() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARINT, 0)
}

func (s *VarintTypeContext) VARINT16() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARINT16, 0)
}

func (s *VarintTypeContext) VARINT32() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARINT32, 0)
}

func (s *VarintTypeContext) VARINT64() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARINT64, 0)
}

func (s *VarintTypeContext) VARSIZE() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARSIZE, 0)
}

func (s *VarintTypeContext) VARUINT() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARUINT, 0)
}

func (s *VarintTypeContext) VARUINT16() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARUINT16, 0)
}

func (s *VarintTypeContext) VARUINT32() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARUINT32, 0)
}

func (s *VarintTypeContext) VARUINT64() antlr.TerminalNode {
	return s.GetToken(ZserioParserVARUINT64, 0)
}

func (s *VarintTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarintTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarintTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterVarintType(s)
	}
}

func (s *VarintTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitVarintType(s)
	}
}

func (s *VarintTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitVarintType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) VarintType() (localctx IVarintTypeContext) {
	localctx = NewVarintTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ZserioParserRULE_varintType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(741)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 86)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 86))) & ((1 << (ZserioParserVARINT - 86)) | (1 << (ZserioParserVARINT16 - 86)) | (1 << (ZserioParserVARINT32 - 86)) | (1 << (ZserioParserVARINT64 - 86)) | (1 << (ZserioParserVARSIZE - 86)) | (1 << (ZserioParserVARUINT - 86)) | (1 << (ZserioParserVARUINT16 - 86)) | (1 << (ZserioParserVARUINT32 - 86)) | (1 << (ZserioParserVARUINT64 - 86)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IFixedBitFieldTypeContext is an interface to support dynamic dispatch.
type IFixedBitFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFixedBitFieldTypeContext differentiates from other interfaces.
	IsFixedBitFieldTypeContext()
}

type FixedBitFieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedBitFieldTypeContext() *FixedBitFieldTypeContext {
	var p = new(FixedBitFieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_fixedBitFieldType
	return p
}

func (*FixedBitFieldTypeContext) IsFixedBitFieldTypeContext() {}

func NewFixedBitFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedBitFieldTypeContext {
	var p = new(FixedBitFieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_fixedBitFieldType

	return p
}

func (s *FixedBitFieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedBitFieldTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZserioParserCOLON, 0)
}

func (s *FixedBitFieldTypeContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZserioParserDECIMAL_LITERAL, 0)
}

func (s *FixedBitFieldTypeContext) BIT_FIELD() antlr.TerminalNode {
	return s.GetToken(ZserioParserBIT_FIELD, 0)
}

func (s *FixedBitFieldTypeContext) INT_FIELD() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT_FIELD, 0)
}

func (s *FixedBitFieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedBitFieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FixedBitFieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFixedBitFieldType(s)
	}
}

func (s *FixedBitFieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFixedBitFieldType(s)
	}
}

func (s *FixedBitFieldTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFixedBitFieldType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FixedBitFieldType() (localctx IFixedBitFieldTypeContext) {
	localctx = NewFixedBitFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ZserioParserRULE_fixedBitFieldType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(743)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZserioParserBIT_FIELD || _la == ZserioParserINT_FIELD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(744)
		p.Match(ZserioParserCOLON)
	}
	{
		p.SetState(745)
		p.Match(ZserioParserDECIMAL_LITERAL)
	}



	return localctx
}


// IDynamicBitFieldTypeContext is an interface to support dynamic dispatch.
type IDynamicBitFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDynamicBitFieldTypeContext differentiates from other interfaces.
	IsDynamicBitFieldTypeContext()
}

type DynamicBitFieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicBitFieldTypeContext() *DynamicBitFieldTypeContext {
	var p = new(DynamicBitFieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_dynamicBitFieldType
	return p
}

func (*DynamicBitFieldTypeContext) IsDynamicBitFieldTypeContext() {}

func NewDynamicBitFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicBitFieldTypeContext {
	var p = new(DynamicBitFieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_dynamicBitFieldType

	return p
}

func (s *DynamicBitFieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DynamicBitFieldTypeContext) BIT_FIELD() antlr.TerminalNode {
	return s.GetToken(ZserioParserBIT_FIELD, 0)
}

func (s *DynamicBitFieldTypeContext) INT_FIELD() antlr.TerminalNode {
	return s.GetToken(ZserioParserINT_FIELD, 0)
}

func (s *DynamicBitFieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicBitFieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DynamicBitFieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterDynamicBitFieldType(s)
	}
}

func (s *DynamicBitFieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitDynamicBitFieldType(s)
	}
}

func (s *DynamicBitFieldTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitDynamicBitFieldType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) DynamicBitFieldType() (localctx IDynamicBitFieldTypeContext) {
	localctx = NewDynamicBitFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, ZserioParserRULE_dynamicBitFieldType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(747)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZserioParserBIT_FIELD || _la == ZserioParserINT_FIELD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IBoolTypeContext is an interface to support dynamic dispatch.
type IBoolTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTypeContext differentiates from other interfaces.
	IsBoolTypeContext()
}

type BoolTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTypeContext() *BoolTypeContext {
	var p = new(BoolTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_boolType
	return p
}

func (*BoolTypeContext) IsBoolTypeContext() {}

func NewBoolTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTypeContext {
	var p = new(BoolTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_boolType

	return p
}

func (s *BoolTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ZserioParserBOOL, 0)
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) BoolType() (localctx IBoolTypeContext) {
	localctx = NewBoolTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, ZserioParserRULE_boolType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(749)
		p.Match(ZserioParserBOOL)
	}



	return localctx
}


// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTypeContext() *StringTypeContext {
	var p = new(StringTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_stringType
	return p
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(ZserioParserSTRING, 0)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterStringType(s)
	}
}

func (s *StringTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitStringType(s)
	}
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) StringType() (localctx IStringTypeContext) {
	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, ZserioParserRULE_stringType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(751)
		p.Match(ZserioParserSTRING)
	}



	return localctx
}


// IFloatTypeContext is an interface to support dynamic dispatch.
type IFloatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatTypeContext differentiates from other interfaces.
	IsFloatTypeContext()
}

type FloatTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatTypeContext() *FloatTypeContext {
	var p = new(FloatTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_floatType
	return p
}

func (*FloatTypeContext) IsFloatTypeContext() {}

func NewFloatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatTypeContext {
	var p = new(FloatTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_floatType

	return p
}

func (s *FloatTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatTypeContext) FLOAT16() antlr.TerminalNode {
	return s.GetToken(ZserioParserFLOAT16, 0)
}

func (s *FloatTypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(ZserioParserFLOAT32, 0)
}

func (s *FloatTypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(ZserioParserFLOAT64, 0)
}

func (s *FloatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FloatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterFloatType(s)
	}
}

func (s *FloatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitFloatType(s)
	}
}

func (s *FloatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitFloatType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) FloatType() (localctx IFloatTypeContext) {
	localctx = NewFloatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, ZserioParserRULE_floatType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(753)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 43)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 43))) & ((1 << (ZserioParserFLOAT16 - 43)) | (1 << (ZserioParserFLOAT32 - 43)) | (1 << (ZserioParserFLOAT64 - 43)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IExternTypeContext is an interface to support dynamic dispatch.
type IExternTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternTypeContext differentiates from other interfaces.
	IsExternTypeContext()
}

type ExternTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternTypeContext() *ExternTypeContext {
	var p = new(ExternTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZserioParserRULE_externType
	return p
}

func (*ExternTypeContext) IsExternTypeContext() {}

func NewExternTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternTypeContext {
	var p = new(ExternTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZserioParserRULE_externType

	return p
}

func (s *ExternTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternTypeContext) EXTERN() antlr.TerminalNode {
	return s.GetToken(ZserioParserEXTERN, 0)
}

func (s *ExternTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExternTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.EnterExternType(s)
	}
}

func (s *ExternTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZserioParserListener); ok {
		listenerT.ExitExternType(s)
	}
}

func (s *ExternTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZserioParserVisitor:
		return t.VisitExternType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZserioParser) ExternType() (localctx IExternTypeContext) {
	localctx = NewExternTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, ZserioParserRULE_externType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(755)
		p.Match(ZserioParserEXTERN)
	}



	return localctx
}


func (p *ZserioParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 52:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZserioParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
			return p.Precpred(p.GetParserRuleContext(), 21)

	case 12:
			return p.Precpred(p.GetParserRuleContext(), 20)

	case 13:
			return p.Precpred(p.GetParserRuleContext(), 19)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

