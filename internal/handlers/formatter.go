package handlers

import (
	"github/internal/models"
	"fmt"
)

type Format struct {
	Events *[]models.GitHubEvent
}

type Formatter interface {
	formatWatchEvent(event *models.GitHubEvent) string
	formatCommentCommitEvent(event *models.GitHubEvent) string
	formatCreateEvent(event *models.GitHubEvent) string
	formatDeleteEvent(event *models.GitHubEvent) string
	formatForkEvent(event *models.GitHubEvent) string
	formatGollumEvent(event *models.GitHubEvent) string
	formatIssueCommentEvent(event *models.GitHubEvent) string
	formatIssueEvent(event *models.GitHubEvent) string
	formatMemberEvent(event *models.GitHubEvent) string
	formatPublicEvent(event *models.GitHubEvent) string
	formatPullReviewEvent(event *models.GitHubEvent) string
	formatPushEvent(event *models.GitHubEvent) string
	formatReleaseEvent(event *models.GitHubEvent) string
	formatSponsorshipEvent(event *models.GitHubEvent) string
	PrintEvents()
}

func (f Format) formatWatchEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("Starred %s", event.Repo.Name)
}

func (f Format) formatCommentCommitEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("Coment Commit on %s and the action was %s", event.Repo.Name, event.Payload.Action)
}

func (f Format) formatCreateEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("Created tag or branch on %s", event.Repo.Name)
}

func (f Format) formatDeleteEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("Tag or branch was deleted on %s", event.Repo.Name)
}

func (f Format) formatForkEvent(event *models.GitHubEvent) string {
	if (event.Payload.Forkee != nil) {
		return fmt.Sprintf("Forked a repo %s", event.Payload.Forkee.Name)
	} 

	return "Forked repo"
}

func (f Format) formatGollumEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("The wiki of %s was updated", event.Repo.Name)
}

func (f Format) formatIssueCommentEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("A comment was %s on %s", event.Payload.Action, event.Repo.Name)
}

func (f Format) formatIssueEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("An issue was %s on %s", event.Payload.Action, event.Repo.Name)
}

func (f Format) formatMemberEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("An user was added on %s", event.Repo.Name)
}

func (f Format) formatPublicEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("A pull request was %s on %s", event.Payload.Action, event.Repo.Name)
}

func (f Format) formatPullReviewEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("A pull request review was updated on %s", event.Repo.Name)
}

func (f Format) formatPushEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("Pushed %d commits to %s", len(*event.Payload.Commits), event.Repo.Name)
}

func (f Format) formatReleaseEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("A release was ocurred on %s", event.Repo.Name)
}

func (f Format) formatSponsorshipEvent(event *models.GitHubEvent) string {
	return fmt.Sprintf("A release was ocurred on %s", event.Repo.Name)
}

func (f Format) PrintEvents(){
	formatterMap := map[string]func(*models.GitHubEvent) string{
		"CommentCommitEvent":         f.formatCommentCommitEvent,
		"WatchEvent":                 f.formatWatchEvent,
		"CreateEvent":                f.formatCreateEvent,
		"DeleteEvent":                f.formatDeleteEvent,
		"GollumEvent":                f.formatGollumEvent,
		"IssueCommentEvent":         f.formatIssueCommentEvent,
		"IssuesEvent":               f.formatIssueEvent,
		"MemberEvent":               f.formatMemberEvent,
		"PublicEvent":               f.formatPublicEvent,
		"PullRequestEvent":          f.formatPublicEvent,           // usa el mismo
		"PullRequestReviewEvent":    f.formatPullReviewEvent,
		"PullRequestReviewCommentEvent": f.formatPullReviewEvent,
		"PullRequestReviewThreadEvent":  f.formatPullReviewEvent,
		"PushEvent":                 f.formatPushEvent,
		"ReleaseEvent":              f.formatReleaseEvent,
		"SponsorshipEvent":          f.formatSponsorshipEvent,
	}

	for _, event := range *f.Events {
		if formatFn, ok := formatterMap[event.Type]; ok {
			fmt.Println(formatFn(&event))
		} else {
			fmt.Println("Unknown event", event.Type)
		}
	}
}

