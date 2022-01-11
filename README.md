# current-next
a small cli to keep track of my current and next task with taskwarrior


View the current task
$ current
17 | back nvim config on github

View the next task
$ next 
18 | read tview docs

Set the current task to something else
$ workon 20
current task set to 20

Set the next task
$ setnext 21 
next task set to 21

Complete a task (Task XX done)
$ complete 21

Completing a task in taskwarrior moves the next task to current automatically.

Data is held in ```~/.current-next/data```
