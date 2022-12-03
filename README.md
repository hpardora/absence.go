# Absencer Client
Absencer has a cron inside that will execute every day at 7:00am and notify you to telegram

Absencer takes in consideration:
- Days that are holiday for your company
- Days that have absences requested by you that are not holiday type

## Configuration:
You need to add the following env vars:
- ABSENCE_PATH: /path/to/absence.yaml (config file)

### Config File content
Config file options example
```yaml
---
# Absence refered
absence_id: absence_api_id
absence_key: absence_api_secred

# working hours
start_hour: '08:30'
end_hour: '18:00'
# random minutes to add to the start and end hours
random_minutes: 10

# days that you will work
working_days: [1,2,3,4,5]

# telegram section
telegram_enabled: true
telegram_api_token: telegra_bot_id:telegram_bot_key
telegram_channel_id: telegram_channel_id
telegram_channel_name: telegram_channel_name
``` 

## TODO
- cron execution time can be defined as variable
- TimeZone can be selected
- Check that clock-out has a valid id and don't explode when there are an error on clock-in
