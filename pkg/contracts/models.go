package contracts

import (
	"time"
)

type ContractSigned struct {
	DocumentID string `json:"documentId,omitempty"`
}

type ContractSignedEvent struct {
	ActingUserEmail     string `json:"actingUserEmail,omitempty"`
	ActingUserId        string `json:"actingUserId,omitempty"`
	ActingUserIpAddress string `json:"actingUserIpAddress,omitempty"`
	ActionType          string `json:"actionType,omitempty"`
	Agreement           struct {
		CreatedDate   *time.Time `json:"createdDate,omitempty"`
		DocumentsInfo struct {
			Documents []struct {
				ID       string `json:"id,omitempty"`
				MimeType string `json:"mimeType,omitempty"`
				Name     string `json:"name,omitempty"`
				NumPages int    `json:"numPages,omitempty"`
			} `json:"documents,omitempty"`
		} `json:"documentsInfo,omitempty"`
		DocumentVisibilityEnabled bool `json:"documentVisibilityEnabled,omitempty"`
		ExternalID                struct {
			ID string `json:"id,omitempty"`
		} `json:"externalId,omitempty"`
		FirstReminderDelay  int    `json:"firstReminderDelay,omitempty"`
		ID                  string `json:"id,omitempty"`
		Locale              string `json:"locale,omitempty"`
		Name                string `json:"name,omitempty"`
		ParticipantSetsInfo struct {
			ParticipantSets []struct {
				ID          string `json:"id,omitempty"`
				MemberInfos []struct {
					Company string `json:"company,omitempty"`
					Email   string `json:"email,omitempty"`
					ID      string `json:"id,omitempty"`
					Name    string `json:"name,omitempty"`
					Status  string `json:"status,omitempty"`
				} `json:"memberInfos,omitempty"`
				Order  int    `json:"order,omitempty"`
				Role   string `json:"role,omitempty"`
				Status string `json:"status,omitempty"`
			} `json:"participantSets,omitempty"`
		} `json:"participantSetsInfo,omitempty"`
		ReminderFrequency string `json:"reminderFrequency,omitempty"`
		SenderEmail       string `json:"senderEmail,omitempty"`
		SignatureType     string `json:"signatureType,omitempty"`
		Status            string `json:"status,omitempty"`
	} `json:"agreement,omitempty"`
	Event                              string     `json:"event,omitempty"`
	EventDate                          *time.Time `json:"eventDate,omitempty"`
	EventResourceType                  string     `json:"eventResourceType,omitempty"`
	InitiatingUserEmail                string     `json:"initiatingUserEmail,omitempty"`
	InitiatingUserId                   string     `json:"initiatingUserId,omitempty"`
	ParticipantRole                    string     `json:"participantRole,omitempty"`
	ParticipantUserEmail               string     `json:"participantUserEmail,omitempty"`
	ParticipantUserId                  string     `json:"participantUserId,omitempty"`
	WebhookID                          string     `json:"webhookId,omitempty"`
	WebhookName                        string     `json:"webhookName,omitempty"`
	WebhookNotificationApplicableUsers []struct {
		Email             string `json:"email,omitempty"`
		ID                string `json:"id,omitempty"`
		PayloadApplicable bool   `json:"payloadApplicable,omitempty"`
		Role              string `json:"role,omitempty"`
	} `json:"webhookNotificationApplicableUsers,omitempty"`
	WebhookNotificationID string `json:"webhookNotificationId,omitempty"`
	WebhookScope          string `json:"webhookScope,omitempty"`
	WebhookUrlInfo        struct {
		URL string `json:"url,omitempty"`
	} `json:"webhookUrlInfo,omitempty"`
}

type Contract struct {
	AcceptedOnDate    time.Time `json:"acceptedOnDate,omitempty" bson:"acceptedOnDate,omitempty"`
	CancellationTerms string    `json:"cancellationTerms,omitempty" bson:"cancellationTerms,omitempty"`
	ContractValue     float64   `json:"contractValue,omitempty" bson:"contractValue,omitempty"`
	CreatedAt         time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ExternalIDs       []string  `json:"externalIDs,omitempty" bson:"externalIDs,omitempty"`
	ID                string    `json:"contractID,omitempty" bson:"contractID,omitempty"`
	OrderID           string    `json:"orderID,omitempty" bson:"orderID,omitempty"`
	OrderLineIDs      []string  `json:"orderLineIDs,omitempty" bson:"orderLineIDs,omitempty"`
	UpdatedAt         time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
