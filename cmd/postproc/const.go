package main

const (
	time = "time"
	issue = "issue"
	message = "message"
	revokes = "revokes"
	fn = "cr-issues.csv"
	msgCheckingWithin = "Checking if record is within allowed time range"
	msgWithin = "Record is within start and end date"
	msgFuture = "Record has a future start date"
	msgSkipping = "Skipping record since start date is in the future"
	msgWithinNoEnd = "Record is within start date (no end date)"
	msgProcessingAccess = "Processing access record"
	msgUpdatingOktaGroup = "Updating Okta group membership"
	msgFoundOkta = "Found Okta user information"
	msgPostedSlack = "Posted Slack message"
	msgSlackIssue = "Posting issue to slack"
	msgAssemblingJira = "Assembling Jira comment"
	msgHandlingLdap = "Handling user LDAP information"
	msgJiraComment = "Added Jira comment"
	msgAddingJiraComment = "Adding Jira comment"
	msgCheckingJiraReviewers = "Checking Jira issue Reviewers"
	msgPreparedUpdate = "Prepared updated Jira issue"
	msgAddingApprovers = "Adding Jira issue Approvers"
	msgUpdatingReviewers = "Updated Jira issue Reviewers"
	msgHandledWebevent = "Handled Webevent"
	msgApprovers = "Updated Jira issue Approvers"
	msgCreatingTasks = "Creating tasks from systems yml definition"
	msgUpdateStatusGranted = "Updated record's access_status to GRANTED"
	msgRecordWithin = "Found record within time range"
	msgSubtask = "Sub-task created"
	msgExpiry = "Sending stack access expiry notification email"
	msgScaleFTNoUser = "ScaleFT user does not exist"
	msgSkippingCauseSubtask = "Skipping record since SplunkCloud Okta subtask is not closed"
	msgAddUser = "Added user to Okta group"
	msgOutside = "Record is outside of allowed time range"
	msgRemoveScaleFT = "Removed record's scaleft_status"
	msgSkipAutoAddMgr = "Skipping auto-adding of manager due to incomplete user verification"
)
