package main

import "testing"

func TestLineRight(t *testing.T) {
	l, err := getLinePath("R1")
	if err != nil {
		t.Error(err)
	}

	if length := len(l); length != 1 {
		t.Fatalf("expected length 1 for line with R1, but got %d: %v", length, l)
	}

	if l[0] != (coord{x: 1, y: 0}) {
		t.Errorf("expected coord to be right of origin by one, instead got %v", l[0])
	}
}

func TestLineLeft(t *testing.T) {
	l, err := getLinePath("L1")
	if err != nil {
		t.Error(err)
	}

	if length := len(l); length != 1 {
		t.Fatalf("expected length 1 for line with L1, but got %d: %v", length, l)
	}

	if l[0] != (coord{x: -1, y: 0}) {
		t.Errorf("expected coord to be right of origin by one, instead got %v", l[0])
	}
}

func TestLineUp(t *testing.T) {
	l, err := getLinePath("U1")
	if err != nil {
		t.Error(err)
	}

	if length := len(l); length != 1 {
		t.Fatalf("expected length 1 for line with U1, but got %d: %v", length, l)
	}

	if l[0] != (coord{x: 0, y: 1}) {
		t.Errorf("expected coord to be right of origin by one, instead got %v", l[0])
	}
}

func TestVisit(t *testing.T) {
	visitedPoints := make(map[int]map[int]int)

	steps := 1
	visited, _ := visit(coord{x: 0, y: 0}, visitedPoints, steps)
	if visited {
		t.Errorf("Shouldn't have been visited when visitedPoints is empty")
	}

	ys, ok := visitedPoints[0]
	if !ok {
		t.Fatal("Expected visited points at 0 to be present")
	}

	if s, ok := ys[0]; !ok || s != steps {
		t.Errorf("Expected ok=true and steps=1, got ok=%v and steps=%v", ok, s)
	}
}

func TestLineDown(t *testing.T) {
	l, err := getLinePath("D1")
	if err != nil {
		t.Error(err)
	}

	if length := len(l); length != 1 {
		t.Errorf("Expected length 1 for line with D1, but got %d", length)
	}

	if l[0] != (coord{x: 0, y: -1}) {
		t.Errorf("Expected coord to be right of origin by one, instead got %v", l[0])
	}
}

func TestIntersectionsOfShortLines(t *testing.T) {
	intersections, err := findIntersections("R1", "U1,R1,D1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil {
		t.Fatal("intersections was nil")
	}

	if l := len(intersections); l != 1 {
		t.Errorf("Expected one intersection, instead got %d", l)
	}

	if i := intersections[0]; i.coord != (coord{x: 1, y: 0}) {
		t.Errorf("Expected intersection at 1,0 , instead found at %v", i)
	}
}

func TestIntersectionsOfShortLines2(t *testing.T) {
	intersections, err := findIntersections("R2,D2", "U3,R5,D5,L3")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil {
		t.Fatal("intersections was nil")
	}

	if l := len(intersections); l != 1 {
		t.Errorf("Expected one intersection, instead got %d", l)
	}

	if i := intersections[0]; i.coord != (coord{x: 2, y: -2}) {
		t.Errorf("Expected intersection at 2,-2 , instead found at %v", i)
	}
}

func TestIntersectionsOfShortLines3(t *testing.T) {
	intersections, err := findIntersections("U3,R5,D5,L3", "R2,D2")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil {
		t.Fatal("intersections was nil")
	}

	if l := len(intersections); l != 1 {
		t.Errorf("Expected one intersection, instead got %d", l)
	}

	if i := intersections[0]; i.coord != (coord{x: 2, y: -2}) {
		t.Errorf("Expected intersection at 2,-2 , instead found at %v", i)
	}
}

func TestCloseToRightDoesntIntersect(t *testing.T) {
	intersections, err := findIntersections("R2", "U1,R3,D1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections != nil {
		t.Errorf("Expected no intersections, but got %v", intersections)
	}
}

func TestCloseToLeftDoesntIntersect(t *testing.T) {
	intersections, err := findIntersections("L2", "U1,L3,D1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections != nil {
		t.Errorf("Expected no intersections, but got %v", intersections)
	}
}

func TestCloseAboveDoesntIntersect(t *testing.T) {
	intersections, err := findIntersections("U1", "R1,U2,L1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections != nil {
		t.Errorf("Expected no intersections, but got %v", intersections)
	}
}

func TestCloseBelowDoesntIntersect(t *testing.T) {
	intersections, err := findIntersections("D1", "R1,D2,L1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections != nil {
		t.Errorf("Expected no intersections, but got %v", intersections)
	}
}
func TestToRightDoesIntersect(t *testing.T) {
	intersections, err := findIntersections("R2", "U1,R2,D1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil || len(intersections) != 1 {
		t.Errorf("Expected intersections, but got %v", intersections)
	}
}

func TestToLeftDoesIntersect(t *testing.T) {
	intersections, err := findIntersections("L2", "U1,L2,D1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil || len(intersections) != 1 {
		t.Errorf("Expected intersections, but got %v", intersections)
	}
}

func TestAboveDoesIntersect(t *testing.T) {
	intersections, err := findIntersections("U1", "R1,U1,L1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil || len(intersections) != 1 {
		t.Errorf("Expected intersections, but got %v", intersections)
	}
}

func TestBelowDoesIntersect(t *testing.T) {
	intersections, err := findIntersections("D1", "R1,D1,L1")
	if err != nil {
		t.Fatal(err)
	}

	if intersections == nil || len(intersections) != 1 {
		t.Errorf("Expected intersections, but got %v", intersections)
	}
}

func TestFindClosest(t *testing.T) {
	closestIntersection, err := findClosestIntersection("R8,U5,L5,D3", "U7,R6,D4,L4")
	if err != nil {
		t.Fatal(err)
	}

	expected := coord{x: 3, y: 3}
	if closestIntersection.coord != expected {
		t.Errorf("Expected %v but got %v", expected, *closestIntersection)
	}
}

func TestFindClosestDistance(t *testing.T) {
	distance, err := distanceFromStartToClosestIntersection("R8,U5,L5,D3", "U7,R6,D4,L4")
	if err != nil {
		t.Fatal(err)
	}

	if expected := 6; distance != expected {
		t.Errorf("Expected %v but got %v", expected, distance)
	}
}

func TestFindClosestDistance2(t *testing.T) {
	distance, err := distanceFromStartToClosestIntersection("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83")
	if err != nil {
		t.Fatal(err)
	}

	if expected := 159; distance != expected {
		t.Errorf("Expected %v but got %v", expected, distance)
	}
}

func TestFindClosestDistance3(t *testing.T) {
	distance, err := distanceFromStartToClosestIntersection("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	if err != nil {
		t.Fatal(err)
	}

	if expected := 135; distance != expected {
		t.Errorf("Expected %v but got %v", expected, distance)
	}
}

func TestSwapInputsForDistanceSame(t *testing.T) {
	input1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	input2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	distance1, err := distanceFromStartToClosestIntersection(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	distance2, err := distanceFromStartToClosestIntersection(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

	if distance1 != distance2 {
		t.Errorf("Got different distances for swapped inputs: %d, %d", distance1, distance2)
	}
}

func TestSwapInputsForDistanceSame2(t *testing.T) {
	input1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	input2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	distance1, err := distanceFromStartToClosestIntersection(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	distance2, err := distanceFromStartToClosestIntersection(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

	if distance1 != distance2 {
		t.Errorf("Got different distances for swapped inputs: %d, %d", distance1, distance2)
	}
}

func TestSwapInputsForDistanceSame3(t *testing.T) {
	input1 := "R8,U5,L5,D3"
	input2 := "U7,R6,D4,L4"

	distance1, err := distanceFromStartToClosestIntersection(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	distance2, err := distanceFromStartToClosestIntersection(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

	if distance1 != distance2 {
		t.Errorf("Got different distances for swapped inputs: %d, %d", distance1, distance2)
	}
}

func TestSwapInputsForDistanceSame4(t *testing.T) {
	input1 := "R1009,D117,L888,D799,L611,U766,L832,U859,L892,D79,R645,U191,L681,D787,R447,D429,L988,U536,L486,D832,R221,D619,R268,D545,L706,U234,L528,D453,R493,D24,L688,U658,L74,D281,R910,D849,L5,U16,R935,D399,L417,U609,R22,D782,L432,D83,L357,D982,L902,U294,L338,U102,R342,D621,R106,U979,L238,U158,R930,D948,L700,D808,R445,U897,R980,U227,L466,D416,R244,U396,R576,U157,R548,U795,R709,U550,R137,U212,L977,U786,L423,D792,R391,D974,R390,U771,R270,D409,L917,D9,R412,D699,L170,D276,L912,U710,R814,U656,R4,D800,R596,U970,L194,U315,L845,D490,L303,U514,L675,D737,L880,D86,L253,D525,R861,D5,R424,D113,L764,D900,R485,D421,R125,U684,R53,U96,L871,U260,R456,U378,L448,D450,L903,D482,R750,U961,R264,D501,R605,D367,R550,U642,R228,U164,L343,U868,R595,D318,R452,U845,L571,D281,R49,D889,L481,U963,R182,U358,R454,U267,L790,D252,R455,D188,L73,U256,L835,D816,R503,U895,L259,U418,R642,U818,L187,U355,R772,U466,R21,U91,R707,D349,L200,U305,R931,D982,L334,D416,L247,D935,L326,U449,L398,D914,R602,U10,R762,D944,L639,D141,L457,U579,L198,U527,R750,U167,R816,D753,R850,D281,L712,D583,L172,D254,L544,D456,R966,U839,R673,D479,R730,D912,R992,D969,R766,U205,R477,D719,R172,D735,R998,D687,R698,D407,R172,U945,R199,U348,L256,D876,R580,U770,L483,D437,R353,D214,R619,U541,R234,D962,R842,U639,R520,D354,L279,D15,R42,U138,L321,D376,L628,D893,L670,D574,L339,U298,L321,D120,L370,U408,L333,D353,L263,D79,R535,D487,R113,D638,R623,D59,L508,D866,R315,U166,L534,U927,L401,D626,L19,D994,L778,D317,L936,U207,L768,U948,R452,U165,R864,D283,L874"
	input2 := "L995,D93,L293,U447,L793,D605,R497,D155,L542,D570,R113,D779,L510,U367,L71,D980,R237,U290,L983,U49,R745,U182,L922,D174,L189,D629,R315,D203,R533,U72,L981,D848,L616,U654,R445,D864,R526,D668,L678,U378,L740,D840,L202,D429,R136,D998,L116,D554,L893,U759,R617,U942,R999,U582,L220,U447,R895,D13,R217,U743,L865,U950,R91,D381,R662,D518,L798,D637,L213,D93,L231,D185,R704,U581,L268,U773,R405,U862,R796,U73,L891,U553,L952,U450,R778,D868,R329,D669,L182,U378,L933,D83,R574,U807,R785,D278,R139,D362,R8,U546,R651,U241,L462,D309,L261,D307,L85,U701,L913,U271,R814,U723,L777,D256,R417,U814,L461,U652,R198,D747,R914,U520,R806,U956,L771,D229,R984,U685,R663,D812,R650,U214,R839,U574,L10,U66,R644,D371,L917,D819,L73,D236,R277,U611,R390,U723,L129,D496,L552,D451,R584,U105,L805,U165,R179,D372,L405,D702,R14,U332,L893,D419,R342,D146,R907,D672,L316,U257,L903,U919,L942,U771,R879,U624,L280,U150,L320,U220,R590,D242,R744,U291,R562,U418,L898,U66,L564,U495,R837,D555,L739,D780,R409,D122,L426,D857,R937,D600,R428,D592,R727,U917,R256,D680,L422,U630,L14,U240,R617,D664,L961,D554,L302,U925,L376,D187,L700,D31,L762,U397,L554,D217,R679,D683,R680,D572,R54,D164,L940,D523,R140,U52,L506,D638,R331,D415,R389,D884,R410,D62,R691,U665,R889,U864,L663,D690,R487,U811,L190,U780,L758,U267,R155,D344,L133,D137,R93,D229,L729,U878,L889,D603,R288,U890,R251,U531,L249,D995,R863,D257,R655,D311,R874,U356,L833,U151,L741,U246,R694,D899,L48,U915,L900,U757,L861,U402,R971,U537,R460,D844,R54,U956,L151,U74,R892,U248,R677,D881,R99,D931,R427"

	distance1, err := distanceFromStartToClosestIntersection(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	distance2, err := distanceFromStartToClosestIntersection(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

	if distance1 != distance2 {
		t.Errorf("Got different distances for swapped inputs: %d, %d", distance1, distance2)
	}

	if distance1 != 721 && distance2 != 721 {
		t.Errorf("Part1's answer not found")
	}
}

func TestSwapInputsForClosestIntersection(t *testing.T) {
	input1 := "R1009,D117,L888,D799,L611,U766,L832,U859,L892,D79,R645,U191,L681,D787,R447,D429,L988,U536,L486,D832,R221,D619,R268,D545,L706,U234,L528,D453,R493,D24,L688,U658,L74,D281,R910,D849,L5,U16,R935,D399,L417,U609,R22,D782,L432,D83,L357,D982,L902,U294,L338,U102,R342,D621,R106,U979,L238,U158,R930,D948,L700,D808,R445,U897,R980,U227,L466,D416,R244,U396,R576,U157,R548,U795,R709,U550,R137,U212,L977,U786,L423,D792,R391,D974,R390,U771,R270,D409,L917,D9,R412,D699,L170,D276,L912,U710,R814,U656,R4,D800,R596,U970,L194,U315,L845,D490,L303,U514,L675,D737,L880,D86,L253,D525,R861,D5,R424,D113,L764,D900,R485,D421,R125,U684,R53,U96,L871,U260,R456,U378,L448,D450,L903,D482,R750,U961,R264,D501,R605,D367,R550,U642,R228,U164,L343,U868,R595,D318,R452,U845,L571,D281,R49,D889,L481,U963,R182,U358,R454,U267,L790,D252,R455,D188,L73,U256,L835,D816,R503,U895,L259,U418,R642,U818,L187,U355,R772,U466,R21,U91,R707,D349,L200,U305,R931,D982,L334,D416,L247,D935,L326,U449,L398,D914,R602,U10,R762,D944,L639,D141,L457,U579,L198,U527,R750,U167,R816,D753,R850,D281,L712,D583,L172,D254,L544,D456,R966,U839,R673,D479,R730,D912,R992,D969,R766,U205,R477,D719,R172,D735,R998,D687,R698,D407,R172,U945,R199,U348,L256,D876,R580,U770,L483,D437,R353,D214,R619,U541,R234,D962,R842,U639,R520,D354,L279,D15,R42,U138,L321,D376,L628,D893,L670,D574,L339,U298,L321,D120,L370,U408,L333,D353,L263,D79,R535,D487,R113,D638,R623,D59,L508,D866,R315,U166,L534,U927,L401,D626,L19,D994,L778,D317,L936,U207,L768,U948,R452,U165,R864,D283,L874"
	input2 := "L995,D93,L293,U447,L793,D605,R497,D155,L542,D570,R113,D779,L510,U367,L71,D980,R237,U290,L983,U49,R745,U182,L922,D174,L189,D629,R315,D203,R533,U72,L981,D848,L616,U654,R445,D864,R526,D668,L678,U378,L740,D840,L202,D429,R136,D998,L116,D554,L893,U759,R617,U942,R999,U582,L220,U447,R895,D13,R217,U743,L865,U950,R91,D381,R662,D518,L798,D637,L213,D93,L231,D185,R704,U581,L268,U773,R405,U862,R796,U73,L891,U553,L952,U450,R778,D868,R329,D669,L182,U378,L933,D83,R574,U807,R785,D278,R139,D362,R8,U546,R651,U241,L462,D309,L261,D307,L85,U701,L913,U271,R814,U723,L777,D256,R417,U814,L461,U652,R198,D747,R914,U520,R806,U956,L771,D229,R984,U685,R663,D812,R650,U214,R839,U574,L10,U66,R644,D371,L917,D819,L73,D236,R277,U611,R390,U723,L129,D496,L552,D451,R584,U105,L805,U165,R179,D372,L405,D702,R14,U332,L893,D419,R342,D146,R907,D672,L316,U257,L903,U919,L942,U771,R879,U624,L280,U150,L320,U220,R590,D242,R744,U291,R562,U418,L898,U66,L564,U495,R837,D555,L739,D780,R409,D122,L426,D857,R937,D600,R428,D592,R727,U917,R256,D680,L422,U630,L14,U240,R617,D664,L961,D554,L302,U925,L376,D187,L700,D31,L762,U397,L554,D217,R679,D683,R680,D572,R54,D164,L940,D523,R140,U52,L506,D638,R331,D415,R389,D884,R410,D62,R691,U665,R889,U864,L663,D690,R487,U811,L190,U780,L758,U267,R155,D344,L133,D137,R93,D229,L729,U878,L889,D603,R288,U890,R251,U531,L249,D995,R863,D257,R655,D311,R874,U356,L833,U151,L741,U246,R694,D899,L48,U915,L900,U757,L861,U402,R971,U537,R460,D844,R54,U956,L151,U74,R892,U248,R677,D881,R99,D931,R427"

	c1, err := findClosestIntersection(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	c2, err := findClosestIntersection(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

	if c1 != c2 {
		t.Errorf("Expected same closest intersection for swapped inputs, but got %v and %v", c1, c2)
	}
}

func TestSwapInputsForFindIntersections(t *testing.T) {
	input1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	input2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	coords1, err := findIntersections(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	coords2, err := findIntersections(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

outer:
	for _, c1 := range coords1 {
		for _, c2 := range coords2 {
			if c1 == c2 {
				continue outer
			}
		}

		// First difference found
		t.Fatalf("Expected same intersections for swapped inputs, but got %v and %v", coords1, coords2)
	}
}

func TestSwapInputsForFindIntersections2(t *testing.T) {
	input1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	input2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	coords1, err := findIntersections(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	coords2, err := findIntersections(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

outer:
	for _, c1 := range coords1 {
		for _, c2 := range coords2 {
			if c1 == c2 {
				continue outer
			}
		}

		// First difference found
		t.Fatalf("Expected same intersections for swapped inputs, but got %v and %v", coords1, coords2)
	}
}

func TestSwapInputsForFindIntersections3(t *testing.T) {
	input1 := "R1009,D117,L888,D799,L611,U766,L832,U859,L892,D79,R645,U191,L681,D787,R447,D429,L988,U536,L486,D832,R221,D619,R268,D545,L706,U234,L528,D453,R493,D24,L688,U658,L74,D281,R910,D849,L5,U16,R935,D399,L417,U609,R22,D782,L432,D83,L357,D982,L902,U294,L338,U102,R342,D621,R106,U979,L238,U158,R930,D948,L700,D808,R445,U897,R980,U227,L466,D416,R244,U396,R576,U157,R548,U795,R709,U550,R137,U212,L977,U786,L423,D792,R391,D974,R390,U771,R270,D409,L917,D9,R412,D699,L170,D276,L912,U710,R814,U656,R4,D800,R596,U970,L194,U315,L845,D490,L303,U514,L675,D737,L880,D86,L253,D525,R861,D5,R424,D113,L764,D900,R485,D421,R125,U684,R53,U96,L871,U260,R456,U378,L448,D450,L903,D482,R750,U961,R264,D501,R605,D367,R550,U642,R228,U164,L343,U868,R595,D318,R452,U845,L571,D281,R49,D889,L481,U963,R182,U358,R454,U267,L790,D252,R455,D188,L73,U256,L835,D816,R503,U895,L259,U418,R642,U818,L187,U355,R772,U466,R21,U91,R707,D349,L200,U305,R931,D982,L334,D416,L247,D935,L326,U449,L398,D914,R602,U10,R762,D944,L639,D141,L457,U579,L198,U527,R750,U167,R816,D753,R850,D281,L712,D583,L172,D254,L544,D456,R966,U839,R673,D479,R730,D912,R992,D969,R766,U205,R477,D719,R172,D735,R998,D687,R698,D407,R172,U945,R199,U348,L256,D876,R580,U770,L483,D437,R353,D214,R619,U541,R234,D962,R842,U639,R520,D354,L279,D15,R42,U138,L321,D376,L628,D893,L670,D574,L339,U298,L321,D120,L370,U408,L333,D353,L263,D79,R535,D487,R113,D638,R623,D59,L508,D866,R315,U166,L534,U927,L401,D626,L19,D994,L778,D317,L936,U207,L768,U948,R452,U165,R864,D283,L874"
	input2 := "L995,D93,L293,U447,L793,D605,R497,D155,L542,D570,R113,D779,L510,U367,L71,D980,R237,U290,L983,U49,R745,U182,L922,D174,L189,D629,R315,D203,R533,U72,L981,D848,L616,U654,R445,D864,R526,D668,L678,U378,L740,D840,L202,D429,R136,D998,L116,D554,L893,U759,R617,U942,R999,U582,L220,U447,R895,D13,R217,U743,L865,U950,R91,D381,R662,D518,L798,D637,L213,D93,L231,D185,R704,U581,L268,U773,R405,U862,R796,U73,L891,U553,L952,U450,R778,D868,R329,D669,L182,U378,L933,D83,R574,U807,R785,D278,R139,D362,R8,U546,R651,U241,L462,D309,L261,D307,L85,U701,L913,U271,R814,U723,L777,D256,R417,U814,L461,U652,R198,D747,R914,U520,R806,U956,L771,D229,R984,U685,R663,D812,R650,U214,R839,U574,L10,U66,R644,D371,L917,D819,L73,D236,R277,U611,R390,U723,L129,D496,L552,D451,R584,U105,L805,U165,R179,D372,L405,D702,R14,U332,L893,D419,R342,D146,R907,D672,L316,U257,L903,U919,L942,U771,R879,U624,L280,U150,L320,U220,R590,D242,R744,U291,R562,U418,L898,U66,L564,U495,R837,D555,L739,D780,R409,D122,L426,D857,R937,D600,R428,D592,R727,U917,R256,D680,L422,U630,L14,U240,R617,D664,L961,D554,L302,U925,L376,D187,L700,D31,L762,U397,L554,D217,R679,D683,R680,D572,R54,D164,L940,D523,R140,U52,L506,D638,R331,D415,R389,D884,R410,D62,R691,U665,R889,U864,L663,D690,R487,U811,L190,U780,L758,U267,R155,D344,L133,D137,R93,D229,L729,U878,L889,D603,R288,U890,R251,U531,L249,D995,R863,D257,R655,D311,R874,U356,L833,U151,L741,U246,R694,D899,L48,U915,L900,U757,L861,U402,R971,U537,R460,D844,R54,U956,L151,U74,R892,U248,R677,D881,R99,D931,R427"

	coords1, err := findIntersections(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	coords2, err := findIntersections(input2, input1)
	if err != nil {
		t.Fatal(err)
	}

outer:
	for _, c1 := range coords1 {
		for _, c2 := range coords2 {
			if c1 == c2 {
				continue outer
			}
		}

		// First difference found
		t.Fatalf("Expected same intersections for swapped inputs, first difference found %v", c1)
	}
}
func TestFindIntersectionsSteps(t *testing.T) {
	input1 := "R8,U5,L5,D3"
	input2 := "U7,R6,D4,L4"

	intersections, err := findIntersections(input1, input2)
	if err != nil {
		t.Fatal(err)
	}

	if len(intersections) != 2 {
		t.Fatalf("Expected one intersection and instead got %d", len(intersections))
	}

	if steps, expected1, expected2 := intersections[0].steps, 40, 30; steps != expected1 && steps != expected2 {
		t.Errorf("Expected number of stpes to be %d or %d but instead got %d", expected1, expected2, steps)
	}

	if steps, expected1, expected2 := intersections[1].steps, 40, 30; steps != expected1 && steps != expected2 {
		t.Errorf("Expected number of stpes to be %d or %d but instead got %d", expected1, expected2, steps)
	}
}
