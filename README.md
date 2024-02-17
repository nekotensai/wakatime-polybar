# wakatime-polybar
Small go project to fetch time from Wakatime to track into polybar.
## Install
- Compile main.go
- Insert Wakatime API-key into config.json
- Insert polybar module in polybar config
Example of polybar module:
```
[module/wakatime-polybar]
type = custom/script
exec = ~/wakatime-polybar/main
interval = 60
```
