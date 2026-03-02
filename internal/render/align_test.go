package render

import (
	"reflect"
	"testing"
)

func TestCalculatePadding(t *testing.T) {
	const terminalWidth = 80
	const artWidth = 20

	if got := calculatePadding("center", artWidth, terminalWidth); got != 30 {
		t.Fatalf("center padding = %d, want 30", got)
	}
	if got := calculatePadding("right", artWidth, terminalWidth); got != 60 {
		t.Fatalf("right padding = %d, want 60", got)
	}
	if got := calculatePadding("left", artWidth, terminalWidth); got != 0 {
		t.Fatalf("left padding = %d, want 0", got)
	}
}

func TestApplyAlign_Right(t *testing.T) {
	lines := []string{"AA", "BB"}
	got := applyAlign(lines, "right", 6)
	want := []string{"    AA", "    BB"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("applyAlign(right) = %#v, want %#v", got, want)
	}
}

func TestApplyAlign_Left(t *testing.T) {
	lines := []string{"AA", "BB"}
	got := applyAlign(lines, "left", 80)

	if !reflect.DeepEqual(got, lines) {
		t.Fatalf("applyAlign(left) = %#v, want %#v", got, lines)
	}
}

func TestApplyAlign_Justify(t *testing.T) {
	lines := []string{"A  B  C"}
	got := applyAlign(lines, "justify", 11)
	want := []string{"A    B    C"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("applyAlign(justify) = %#v, want %#v", got, want)
	}
}

