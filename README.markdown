# iCal-Filter

Reads iCal events and prints them unless they contain a given string. I use this in a pipe to filter unwanted events:

```command
# Fetch the NASA calendar and remove all SpaceX events
$ curl https://www.nasa.gov/templateimages/redesign/calendar/iCal/nasa_calendar.ics | ical-filter --skip SpaceX
```
