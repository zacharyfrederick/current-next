# current-next
a small cli to keep track of my current and next task with taskwarrior


View the current task
$ current
17 | back nvim config on github

View the next task
$ next 
18 | read tview docs

Set the current task to something else
$ current 20 -s
current task set to 20

Set the next task
$ next 21 -s 
next task set to 21

Complete a task (Task XX done)
$ finish 21

Completing a task in taskwarrior moves the next task to current automatically.

Data is held in ```~/.current-next/data```

# TODO
1. Move code from individual binaries into library
2. Update finish command to move the current contents of next into current
3. remove setnext and workon and move functionality into current and next with flags
