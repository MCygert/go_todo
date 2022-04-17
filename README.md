# Project for saving TO-DO tasks

## How it works
 *Work in progress*
## Simple project made for practicing go
Main functionality is saving task which I want to do in next 24h. 
It could be simple web service but I wanted to fiddle with architecture, so I made 2 different data input sources

* CLI 
* Web

I am thinking about adding some backup db where I could dump expired tasks after some time.
It would increase complexity of this project.

## Technologies
* DB -> mongo 
* CRON -> [robfig/crong](https://github.com/robfig/cron)
* Rest is made with go stdlib

## Future plans for this project
 - [x] Cron for marking expired tasks
 - [ ] Enhance CLI so it can do more commands
 - [ ] Add marking tasks as done (manually)
 - [ ] Add listing all tasks
 - [ ] Sort tasks chronologically
 - [ ] Add notifying if task expires
