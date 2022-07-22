/*
 * Incidents and Alerts API
 *
 *  ## Incidents  An incident is the combination of the alert event and clear event for a detector rule.  Detectors monitor metric time series; in the context of a detector, these are known as a **signal**. When a condition defined in a  **detector rule** matches the signal, the detector does the following:    * Triggers an **alert** that you can view in a number of places throughout SignalFx   * Generates an **event** that you can view in the web UI. SignalFx also stores the     event.   * Sends one or more **notifications**, so that users learn about the alert     even if they're not looking at the web UI.  When the signal no longer matches the condition in the detector rule, the detector generates a **clear event** and sends out a second notification. Together, the alert and clear event for a detector rule constitute an  incident.  ## Alert muting rules  Alert muting rules (known as muting rules in the web UI) stop a detector from sending notifications, based on properties you specify in the rule. When you mute a notification, the detector still triggers alerts and  generates events, and you can see these in their respective locations in the web UI. However, the detector doesn't send notifications to recipients who would normally receive them.  Use alert muting rules to mute certain notifications when you want to schedule server downtime or test new code or configurations. You can set alert muting rules to last for a specific time period or to last indefinitely.  In some cases, SignalFx may send notifications during an alert muting period. To learn more, see the section  [Considerations for alert muting](https://developers.signalfx.com/detectors_events_alerts/detectors_overview.html#_considerations_for_alert_muting) in the API concepts guide.  <div  style=\"   position: relative;    padding: 0.75rem 1.25rem;    margin-bottom: 1rem;    border: 1px solid transparent;    border-radius: 0.25rem;    color: #0c5460;    background-color: #d1ecf1;    border-color: #bee5eb; \"> <strong>NOTE:</strong> Although the Detectors API only works with detectors you create in the API (<strong>v2</strong> detectors, the <code>/incident</code> and <code>/alertmuting</code> endpoints work with <em>all</em> events and incidents, regardless of which detector version created them. This means that you can use the endpoints to work with incidents and alert muting rules for detectors you create in the web UI. </div>
 *
 * API version: 3.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package alertmuting

// Properties of an alert muting rule, in the form of a JSON object. **NOTE:** You can't create or update properties marked read-only. You receive read-only properties in response bodies for the following:    * **GET** `/alertmuting`   * **POST** `/alertmuting`   * **GET** `/alertmuting/{id}`   * **PUT** `/alertmuting/{id}`
type AlertMutingRule struct {
	// The time the alerting muting rule was created, in Unix time UTC-relative.  **This property is read-only; it's always set by the system.**
	Created int64 `json:"created,omitempty"`
	// SignalFx user ID of the alert muting rule creator, in the form of a JSON string. **This property is read-only; it's always set by the system.**
	Creator string `json:"creator,omitempty"`
	// Description of the rule. **read/write**
	Description string `json:"description,omitempty"`
	// List of alert muting filters for this rule, in the form of a JSON array of alert muting filter objects. Each object is a set of conditions for an alert muting rule. Each object property (name-value pair)  specifies a dimension or custom property to match to alert events. **read/write**
	Filters []*AlertMutingRuleFilter `json:"filters,omitempty"`
	// SignalFx-assigned ID of an alert muting rule, in the form of a JSON string. **This property is read-only; it's always set by the system.**
	Id string `json:"id,omitempty"`
	// The last time the alert muting rule was last updated, in Unix time  UTC-relative. **This property is read-only; it's always set by the system.**
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// The SignalFx-assigned user ID of the last user who updated the alert muting rule, in the form of a JSON string. If the system made the  last update, the value is \"AAAAAAAAAA\". **This property is read-only;  it's always set by the system.**
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// Toggle if alerts should be sent once the muting period is over.  If not specified, defaults to false
	SendAlertsOnceMutingPeriodHasEnded bool `json:"sendAlertsOnceMutingPeriodHasEnded"`
	// Starting time of an alert muting rule, in Unix time format UTC. If not specified, defaults to the current time. **read/write**.
	StartTime int64 `json:"startTime,omitempty"`
	// Stop time of an alert muting rule, in Unix time format UTC. If set to 0, detectors that match this rule are muted indefinitely. The default value is 0. **read/write**
	StopTime int64 `json:"stopTime,omitempty"`
}
