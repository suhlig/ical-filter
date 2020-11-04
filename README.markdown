# iCal-Filter

_Work in Progress_

Reads iCal events and prints them unless they contain a given string. I use this in a pipe to filter unwanted events:

```command
$ curl -L https://pagerduty.com/private/deadbeef/feed/EXAMPLE | ical-filter --skip "SPAM" > filtered.ical
```
