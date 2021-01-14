package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var result *widget.Entry = widget.NewMultiLineEntry()
var activeButton *widget.Button = widget.NewButton("Active", active)
var inactiveButton *widget.Button = widget.NewButton("Inactive", inactive)
var checkRuleExeptionsButton *widget.Button = widget.NewButton("Check Rule Exceptions", checkRuleExceptions)
var roleGrantedButton *widget.Button = widget.NewButton("Role Granted", role)
var followUpButton *widget.Button = widget.NewButton("Follow Up", followup)
var followUp2Button *widget.Button = widget.NewButton("More Follow Up", followup2)
var followUpScaleFTButton *widget.Button = widget.NewButton("ScaleFT User DNE", followupscale)
var selectRule *widget.Select
var ruleList []string

func showState(description string, sRule stateRule) {
	jiras := description + "\n"
	count := 0
	for k, v := range messagesByIssue {
		if evaluateRule(v, sRule) {

			jiras = jiras + fmt.Sprintf("%s, ", k)
			count++
			if count & 7 == 0 {
				jiras = jiras + "\n"
			}
		}
	}
	result.SetText(jiras)
}

func showState5() {
	jiras := "follow up\n"
	for k, v := range messagesByIssue {
		if evaluateRule(v, isMsg5) {
			jiras = jiras + fmt.Sprintf("%s: %s\n", k, v[5])
		}
	}
	result.SetText(jiras)
}

func showState5n6() {
	jiras := "More follow up\n"
	for k, v := range messagesByIssue {
		if evaluateRule(v, isMsg5and6) {
			jiras = jiras + fmt.Sprintf("%s: %s\n%s\n\n", k, v[5], v[6])
		}
	}
	result.SetText(jiras)
}


func active() {
	showState("active", isActive)
}

func inactive() {
	showState("inactive", isNotYetActive)}

func role() {
	showState("grant role", isRoleGranted)
}

func followup() {
	showState5()
}

func followup2() {
	showState5n6()
}

func followupscale() {
	showState("ScaleFT user does not exist", isUpdateScaleFTUserDNE)
}

func checkRuleExceptions() {
	jiras := "no matching rule\n"
	var v []string

	for _, v = range messagesByIssue {
		missing := true
		for _,nxtRule := range allRules.rules {
			if evaluateRule(v, nxtRule) {
				missing = false
				break
			}
		}
		if missing {
			jiras = jiras + fmt.Sprintf("\n")
			result.SetText(jiras)
			jiras = jiras + fmt.Sprintf("Next item has %i messages\n", len(v))
			for i := 1; i < len(v); i++ {
				jiras = jiras + fmt.Sprintf("|%s|\n", v[i])
			}
		}
	}
	result.SetText(jiras)
}


func checkRules() {
	for k, v := range messagesByIssue {
		if evaluateRule(v, isActive) {
			fmt.Printf("%s active\n", k)
		} else if evaluateRule(v, isNotYetActive) {
			fmt.Printf("%s not yet active\n", k)
		} else if evaluateRule(v, isJustExpired) {
			fmt.Printf("%s just expired\n", k)
		} else if evaluateRule(v, isRoleGranted) {
			fmt.Printf("%s role granted\n", k)
		} else if evaluateRule(v, isMsg5) {
			fmt.Printf("%s %s\n", k, v[5])
		} else if evaluateRule(v, isMsg5and6) {
			fmt.Printf("%s %s: %s\n", k, v[5], v[6])
		} else if evaluateRule(v, isSkippingOpenSubtasks) {
			fmt.Printf("%s skipping, open subtasks\n", k)
		} else if evaluateRule(v, isSubtaskCreated) {
			fmt.Printf("%s subtask created\n", k)
		} else if evaluateRule(v, isUpdateApprovers) {
			fmt.Printf("%s Updated approvers\n", k)
		} else if evaluateRule(v, isUpdateReviewersAndApprovers) {
			fmt.Printf("%s Updated approvers and reviewers\n", k)
		} else if evaluateRule(v, isUpdateScaleFTUserDNE) {
			fmt.Printf("%s ScaleFT user does not exist\n", k)
		} else if evaluateRule(v, isIncompleteUserVerififation) {
			fmt.Printf("%s Incomplete user verification\n", k)
		} else if evaluateRule(v, isRemoveScaleFT) {
			fmt.Printf("%s Remove record's ScaleFT status\n", k)
		} else {
			fmt.Printf("*******************\n%s \n%s \n\n", k, v[0])
			for i := 1; i < len(v); i++ {
				fmt.Printf("%s \n\n", v[i])
			}
			fmt.Printf("*******************\n")
		}
	}
}

func ruleSelected(selection string) {
	for i := 0; i < len(allRules.rules); i++ {
		if allRules.rules[i].name == selection {
			showState(selection, allRules.rules[i])
		}
	}
}

func setupUi() {
	ruleList = make([]string, len(allRules.rules))
	for i := 0; i < len(allRules.rules); i++ {
		ruleList[i] = allRules.rules[i].name
	}
	selectRule = widget.NewSelect(ruleList, ruleSelected)

	a := app.New()
	w := a.NewWindow("Jira Sorter")

	ui := widget.NewVBox(
		widget.NewHBox(followUpButton, followUp2Button, followUpScaleFTButton),
		widget.NewHBox(selectRule, checkRuleExeptionsButton, activeButton, inactiveButton, roleGrantedButton),
		widget.NewHBox(result),
	)
	w.SetContent(ui)

	w.ShowAndRun()
}
