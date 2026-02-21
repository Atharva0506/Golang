package main

import (
	"testing"
)

// To run this test:
// 1. CD into this '01_Beginner_Test' directory
// 2. Run 'go test -v'

func TestPlayerStruct(t *testing.T) {
	// Task 1 Test
	p := Player{Name: "Gopher", Health: 100, Score: 0, IsAlive: true}

	if p.Name != "Gopher" {
		t.Errorf("Expected player name 'Gopher', got %v", p.Name)
	}

	p.TakeDamage(30)
	if p.Health != 70 {
		t.Errorf("Expected health 70, got %v", p.Health)
	}

	p.TakeDamage(100)
	if p.Health != 0 || p.IsAlive != false {
		t.Errorf("Expected health 0 and IsAlive false, got Health %v, IsAlive %v", p.Health, p.IsAlive)
	}
}

func TestClassifyScores(t *testing.T) {
	// Task 3 Test
	scores := []int{10, 45, 60, 75, 80, 99, 100}

	result := ClassifyScores(scores)

	if result["Low"] != 2 {
		t.Errorf("Expected 2 Low scores, got %v", result["Low"])
	}
	if result["Medium"] != 2 {
		t.Errorf("Expected 2 Medium scores, got %v", result["Medium"])
	}
	if result["High"] != 3 {
		t.Errorf("Expected 3 High scores, got %v", result["High"])
	}
}
