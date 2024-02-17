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
![image](https://github.com/nekotensai/wakatime-polybar/assets/49762125/dad135ed-abda-4010-953d-50a6e80a4e01)


