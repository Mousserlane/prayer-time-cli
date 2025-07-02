
# Prayer Time CLI (WIP)

A simple prayer time program... In your CLI!
I made this to try out golang & BubbleTea. And because I want a running clock & a prayer time reminder in my terminal workflow.

<img width="724" alt="Screenshot 2025-06-17 at 17 45 07" src="https://github.com/user-attachments/assets/1832899c-2c04-4c9d-90e5-e50888e328ec" />

This is still a work in progress. Main feature is ready but further enhancement is required.

## Main Feature
- Digital clock
- Gregorian & Hijri date.
  - Every important dates in islam are in Hijri. Having this feature could hopefully remind fellow muslims about these events.
- Today's prayer time from Fajr to Isha
- Shuruq time

## Upcoming Feature
- Adzan
- Monthly / Weekly prayer time
- Hijri calendar menu
  - Highlight important dates

## Installation and Running the App
Run `make install` to install the app.
To execute the app, simply run `prayer-time-cli` 

## Progress Tracker

- [x] Add Current Time
- [x] Add day month year in hijri
- [x] Add Gregorian Date
- [x] Add Prayer times for the date
- [ ] Add Month/Weekly prayer time
- [x] Add Highlight for nearest prayer time
- [x] Add location and prayer time calculation method selection (At the moment it's hardcoded to Jakarta using ID's ministry of religious affair calculation method)
- [x] Init config prompt i.e. select location or timezone, add lat/long for better precision, and pick calculation method
- [x] Save config
- [ ] Fallback values for empty config item e.g. maps predefined lat/long to a timezone if users omit the lat long config.  
