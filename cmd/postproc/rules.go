package main

type matchRule struct {
	msgNumber	int
	msg	string
}

type stateRule struct {
	name string
	rules []aRule
}

type aRule struct {
	msgCount	int
	matchRules []matchRule
}

type ruleBook struct {
	rules	[]stateRule
}

func evaluateRule(v []string, sRules stateRule) bool {
	var wrong bool

	for _, nxtRule := range sRules.rules {
		wrong = false
		if nxtRule.msgCount == len(v) {
			for _, msgRule := range nxtRule.matchRules {
				if v[msgRule.msgNumber] != msgRule.msg {
					wrong = true
					break
				}
			}
			if !wrong {
				return true
			}
		}
	}
	return false
}

var isRoleGranted = stateRule{"role granted",
	[]aRule{aRule{
		18,
			[]matchRule{
				matchRule{msgNumber: 1, msg: msgPostedSlack},
				matchRule{msgNumber: 2, msg: msgAssemblingJira},
				matchRule{msgNumber: 3, msg: msgHandlingLdap},
				matchRule{msgNumber: 4, msg: msgJiraComment},
				matchRule{msgNumber: 5, msg: msgCheckingJiraReviewers},
				matchRule{msgNumber: 6, msg: msgPreparedUpdate},
				matchRule{msgNumber: 7, msg: msgAddingApprovers},
				matchRule{msgNumber: 8, msg: msgUpdatingReviewers},
				matchRule{msgNumber: 17, msg: msgUpdateStatusGranted},
			},
		},
	},
}

var isActive = stateRule{"active",
	[]aRule{
		aRule{2,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithin},
			},
		},
	},
}

var isNotYetActive = stateRule{"not yet active",
	[]aRule{
		aRule{3,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgFuture},
				matchRule{msgNumber: 2, msg: msgSkipping},
			},
		},
	},
}

var isJustExpired = stateRule{"just expired",
	[]aRule{
		aRule{4,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithin},
				matchRule{msgNumber: 2, msg: msgRecordWithin},
				matchRule{msgNumber: 3, msg: msgExpiry},
			},
		},
	},
}

var isMsg5 = stateRule{"various errors",
	[]aRule{
		aRule{6,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithin},
				matchRule{msgNumber: 3, msg: msgUpdatingOktaGroup},
				matchRule{msgNumber: 4, msg: msgFoundOkta},
			},
		},
		aRule{6,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithinNoEnd},
				matchRule{msgNumber: 3, msg: msgUpdatingOktaGroup},
				matchRule{msgNumber: 4, msg: msgFoundOkta},
			},
		},
		aRule{6,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithinNoEnd},
				matchRule{msgNumber: 4, msg: msgUpdatingOktaGroup},
				matchRule{msgNumber: 2, msg: msgFoundOkta},
			},
		},
	},
}



// Note a copuple of ||, missing from above
//if len(v) == 6 && msgCheckingWithin == v[0] && (msgWithin == v[1] || msgWithinNoEnd == v[1]) && (msgUpdatingOktaGroup == v[3] || msgUpdatingOktaGroup == v[4]) && (msgFoundOkta == v[4] || msgFoundOkta == v[5]){
//fmt.Printf("%s %s\n", k, v[5])

var isMsg5and6 = stateRule{"more errors",
	[]aRule{aRule{7,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 6, msg: msgScaleFTNoUser},
			},
		},
	},
}

var isSkippingOpenSubtasks = stateRule{"skip, open subtasks",
	[]aRule{aRule{3,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithinNoEnd},
				matchRule{msgNumber: 2, msg: msgSkippingCauseSubtask},
			},
		},
	},
}

var isSubtaskCreated = stateRule{"subtask created",
	[]aRule{
		aRule{2,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgSubtask},
				matchRule{msgNumber: 1, msg: msgHandledWebevent},
			},
		},
	},
}

var isUpdateApprovers = stateRule{"update approvers",
	[]aRule{
		aRule{13,
			[]matchRule{
				matchRule{msgNumber: 6, msg: msgHandlingLdap},
				matchRule{msgNumber: 8, msg: msgPreparedUpdate},
				matchRule{msgNumber: 11, msg: msgHandledWebevent},
				matchRule{msgNumber: 12, msg: msgApprovers},
			},
		},
		aRule{12,
			[]matchRule{
				matchRule{msgNumber: 5, msg: msgHandlingLdap},
				matchRule{msgNumber: 9, msg: msgPreparedUpdate},
				matchRule{msgNumber: 10, msg: msgHandledWebevent},
				matchRule{msgNumber: 11, msg: msgApprovers},
			},
		},
	},
}


var isIncompleteUserVerififation = stateRule{"incomplete usere verification",
	[]aRule{
		aRule{13,
		[]matchRule{
				matchRule{msgNumber: 3, msg: msgPostedSlack},
				matchRule{msgNumber: 4, msg: msgAssemblingJira},
				matchRule{msgNumber: 5, msg: msgHandlingLdap},
				matchRule{msgNumber: 6, msg: msgJiraComment},
				matchRule{msgNumber: 7, msg: msgAddingApprovers},
				matchRule{msgNumber: 8, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{13,
			[]matchRule{
				matchRule{msgNumber: 1, msg: msgSlackIssue},
				matchRule{msgNumber: 5, msg: msgAssemblingJira},
				matchRule{msgNumber: 6, msg: msgHandlingLdap},
				matchRule{msgNumber: 7, msg: msgJiraComment},
				matchRule{msgNumber: 8, msg: msgAddingApprovers},
				matchRule{msgNumber: 9, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{13,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgSlackIssue},
				matchRule{msgNumber: 5, msg: msgAssemblingJira},
				matchRule{msgNumber: 6, msg: msgHandlingLdap},
				matchRule{msgNumber: 7, msg: msgJiraComment},
				matchRule{msgNumber: 8, msg: msgAddingApprovers},
				matchRule{msgNumber: 9, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{12,
			[]matchRule{
				matchRule{msgNumber: 9, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{13,
			[]matchRule{
				matchRule{msgNumber: 9, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{12,
			[]matchRule{
				matchRule{msgNumber: 8, msg: msgSkipAutoAddMgr},
			},
		},
		aRule{13,
			[]matchRule{
				matchRule{msgNumber: 8, msg: msgSkipAutoAddMgr},
			},
		},
	},
}

var isUpdateReviewersAndApprovers = stateRule{"update reviewers and approvers",
	[]aRule{
	aRule{11,
		[]matchRule{
				matchRule{msgNumber: 1, msg: msgPostedSlack},
				matchRule{msgNumber: 3, msg: msgHandlingLdap},
				matchRule{msgNumber: 7, msg: msgAddingApprovers},
				matchRule{msgNumber: 8, msg: msgUpdatingReviewers},
				matchRule{msgNumber: 9, msg: msgHandledWebevent},
				matchRule{msgNumber: 10, msg: msgApprovers},
			},
		},
	},
}

var isUpdateScaleFTUserDNE = stateRule{"scaleFT user does not exist",
	[]aRule{
		aRule{3,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 1, msg: msgWithin},
				matchRule{msgNumber: 2, msg: msgScaleFTNoUser},
			},
		},
	},
}

var isRemoveScaleFT = stateRule{"remove scalFT",
	[]aRule{
		aRule{4,
			[]matchRule{
					matchRule{msgNumber: 0, msg: msgCheckingWithin},
					matchRule{msgNumber: 1, msg: msgWithin},
					matchRule{msgNumber: 2, msg: msgOutside},
					matchRule{msgNumber: 3, msg: msgRemoveScaleFT},
			},
		},
		aRule{4,
			[]matchRule{
				matchRule{msgNumber: 0, msg: msgCheckingWithin},
				matchRule{msgNumber: 2, msg: msgWithin},
				matchRule{msgNumber: 1, msg: msgOutside},
				matchRule{msgNumber: 3, msg: msgRemoveScaleFT},
			},
		},
	},
}

var allRules = ruleBook{[]stateRule{isRoleGranted, isActive, isNotYetActive, isJustExpired,
	isMsg5, isMsg5and6, isSkippingOpenSubtasks, isSubtaskCreated, isUpdateApprovers,
	isIncompleteUserVerififation, isUpdateReviewersAndApprovers, isUpdateScaleFTUserDNE,
	isRemoveScaleFT}}