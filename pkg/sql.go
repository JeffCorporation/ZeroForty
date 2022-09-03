package pkg

const sqlCreateProjectTable = `
CREATE TABLE IF NOT EXISTS project
(
    id        INTEGER PRIMARY KEY,
    name      TEXT NOT NULL,
    groupName TEXT NOT NULL,
	uuid      TEXT UNIQUE,
    color     TEXT NOT NULL
);`

const sqlCreateTaskTable = `
CREATE TABLE IF NOT EXISTS task
(
    id        INTEGER PRIMARY KEY,
    title     TEXT NOT NULL,
    projectId INTEGER,
    created   INTEGER,
    completed INTEGER,
    due       INTEGER,
    priority  INTEGER,	
    FOREIGN KEY (projectId) REFERENCES project (id)
);`

const sqlCreateTrackingTable = `
CREATE TABLE IF NOT EXISTS tracking
(
    taskId INTEGER NOT NULL,
    start  INTEGER NOT NULL,
    stop   INTEGER,
    PRIMARY KEY (taskId, start),
    FOREIGN KEY (taskId) REFERENCES task (id)
);`

const sqlSelectProjectNullUUID = `SELECT id FROM project WHERE uuid IS NULL;`
const sqlUpdateProjectUUID = `UPDATE project SET uuid = ? where id = ?;`

const sqlInsertTask = `INSERT INTO task (title, projectId, created, due, completed, priority) VALUES (?,?,?,?,?,?);`
const sqlUpdateTask = `UPDATE task SET title = ?, projectId = ?, due = ?, completed = ?, priority = ? WHERE id = ?;`

const sqlStopTracking = `UPDATE tracking SET stop = ? WHERE stop IS NULL;`
const sqlStartTracking = `INSERT INTO tracking (taskId, start) VALUES (?,?);`

const sqlSelectActiveTask = `
SELECT id, title, IFNULL(projectId,0) as projectId, due, completed, priority
FROM task
JOIN tracking  on task.id = tracking.taskId
WHERE tracking.stop IS NULL
LIMIT 1;`

const sqlSelectTracking = `
SELECT start, stop
FROM tracking WHERE taskId = ? AND start >= ? AND stop <= ?
ORDER BY stop, start;`

const sqlSelectTodoTasks = `
SELECT id, title, IFNULL(projectId,0) as projectId, due, completed, priority, 
       SUM(IFNULL(tr.stop,strftime('%s', 'now') )-tr.start) as timespan
FROM task t
JOIN tracking tr on t.id = tr.taskId
WHERE t.completed IS NULL
GROUP BY id;
`
