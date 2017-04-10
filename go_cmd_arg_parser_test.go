package go_cmd_arg_parser

import "testing"

func TestReturnsAnArrayOfStrings(t *testing.T) {
	args := ParseFromCmdString(`date 1423`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "1423" {
		t.Errorf("Expected the second arg to be '1423', but was '%s'", args[1])
	}
}

func TestRemovesOuterSingleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date '1423'`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "1423" {
		t.Errorf("Expected the second arg to be '1423', but was '%s'", args[1])
	}
}

func TestRemovesOuterDoubleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date "1423"`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "1423" {
		t.Errorf("Expected the second arg to be '1423', but was '%s'", args[1])
	}
}

func TestProperlyHandlesWhitespaceWithinSingleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date -f '%a %b %d %T %Z %Y' "01/01/1900" '+%s'`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "-f" {
		t.Errorf("Expected the second arg to be '-f', but was '%s'", args[1])
	}

	if args[2] != "%a %b %d %T %Z %Y" {
		t.Errorf("Expected the second arg to be '%%a %%b %%d %%T %%Z %%Y', but was '%s'", args[2])
	}

	if args[3] != "01/01/1900" {
		t.Errorf("Expected the second arg to be '01/01/1900', but was '%s'", args[2])
	}

	if args[4] != "+%s" {
		t.Errorf("Expected the second arg to be '+%%s', but was '%s'", args[2])
	}
}

func TestProperlyHandlesWhitespaceWithinDoubleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date -f "%a %b %d %T %Z %Y" "01/01/1900" "+%s"`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "-f" {
		t.Errorf("Expected the second arg to be '-f', but was '%s'", args[1])
	}

	if args[2] != "%a %b %d %T %Z %Y" {
		t.Errorf("Expected the second arg to be '%%a %%b %%d %%T %%Z %%Y', but was '%s'", args[2])
	}

	if args[3] != "01/01/1900" {
		t.Errorf("Expected the second arg to be '01/01/1900', but was '%s'", args[2])
	}

	if args[4] != "+%s" {
		t.Errorf("Expected the second arg to be '+%%s', but was '%s'", args[2])
	}
}

func TestProperlyHandlesEscapedDoubleQuotesWithinDoubleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date -f "%a %b %d \"%T %Z %Y\"" "01/01/1900" "+%s"`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "-f" {
		t.Errorf("Expected the second arg to be '-f', but was '%s'", args[1])
	}

	if args[2] != `%a %b %d \"%T %Z %Y\"` {
		t.Errorf("Expected the second arg to be '%%a %%b %%d \\\"%%T %%Z %%Y\\\"', but was '%s'", args[2])
	}

	if args[3] != "01/01/1900" {
		t.Errorf("Expected the second arg to be '01/01/1900', but was '%s'", args[2])
	}

	if args[4] != "+%s" {
		t.Errorf("Expected the second arg to be '+%%s', but was '%s'", args[2])
	}
}

func TestProperlyHandlesNestedSingleQuotesWithinDoubleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date -f "%a %b %d '%T %Z %Y' -- \"foobar\"" "01/01/1900" "+%s"`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "-f" {
		t.Errorf("Expected the second arg to be '-f', but was '%s'", args[1])
	}

	if args[2] != `%a %b %d '%T %Z %Y' -- \"foobar\"` {
		t.Errorf("Expected the second arg to be '%a %%b %%d '%%T %%Z %%Y' -- \"foobar\"', but was '%s'", args[2])
	}

	if args[3] != "01/01/1900" {
		t.Errorf("Expected the second arg to be '01/01/1900', but was '%s'", args[2])
	}

	if args[4] != "+%s" {
		t.Errorf("Expected the second arg to be '+%%s', but was '%s'", args[2])
	}
}

func TestProperlyHandlesNestedSingleQuotesInNestedDoubleQuotes(t *testing.T) {
	args := ParseFromCmdString(`date -f "\"The title of the book was 'foobar'\" - %a %b %d %T %Z %Y" "01/01/1900" "+%s"`)

	if args[0] != "date" {
		t.Errorf("Expected the first arg to be 'date', but was '%s'", args[0])
	}

	if args[1] != "-f" {
		t.Errorf("Expected the second arg to be '-f', but was '%s'", args[1])
	}

	if args[2] != `\"The title of the book was 'foobar'\" - %a %b %d %T %Z %Y` {
		t.Errorf("Expected the second arg to be '\\\"The title of the book was 'foobar'\\\" - %%a %%b %%d %%T %%Z %%Y', but was '%s'", args[2])
	}

	if args[3] != "01/01/1900" {
		t.Errorf("Expected the second arg to be '01/01/1900', but was '%s'", args[2])
	}

	if args[4] != "+%s" {
		t.Errorf("Expected the second arg to be '+%%s', but was '%s'", args[2])
	}
}
