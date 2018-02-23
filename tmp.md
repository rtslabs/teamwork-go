favorite {
    name,
    taskId,
    taskListid,
    projectId,
    date,
    message,
    hours,
    minutes,
    billable
}

todo {
    name,
    taskId,
    projectId,
    due-date,
    description
}


hours tw stats [summary]
    --date [date/year/month/week/day]
    --favorite

hours tw stats weights
    --used-since [date/year/month/week/day]
    --used-before [date]
    --used-on [date]

hours tw show entries
hours tw show tasks
hours tw show projects
hours tw show taskLists
    --used-since [date/year/month/week/day]
    --used-before [date]
    --used-on [date]
    --assigned
    --tags
    --favorite

hours tw create task
hours tw create taskLists
    --stuff

hours tw log <taskId/favorite>
    -- [ all of favorites ]
    --favorite
    --json opens editor
    --previous
    --git-logs

tw edit <entryId>
    ^^
hours show favorites
    --favorite

hours favorite <name>
    --^^

hours config [<key>[=<value>]]
    --dir
    (key, url, google key)
hours config init
    --dir

hours timer start <name>
    --arrived

hours timer stop <name>
hours timer switch <name>

hours timer time <name>
    --start
    --stop
    --duration
    --stop-all

hours todo add <name>
hours todo complete <name>
    --^^


global flags
    --json

The pattern to follow is APPNAME VERB NOUN --ADJECTIVE. or APPNAME COMMAND ARG --FLAG


src
    main.go
    cmd
        config
            ... .go
        timer
        ...
    configuration
        favorites
    teamwork
    googlecal
    status
    interactive