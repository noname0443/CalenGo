# GroupProject

The project is on an unknown topic.

# Authors
- Ivanov Evgeny
- Ermakov Aleksandr (Ermak)
- Smirnov Petr

# Documentation

# About the project
- We are pleased to present to you another project of ours. This is an interactive calendar. With its help, users will be able to schedule planned events for themselves. They will be able to add new events, as well as edit them if necessary or if something in their plans changes.

# Features
- 1. Reliable registration: a user registered to use the calendar, writing the next event in the schedule, automatically enters it into its database and when restarting and a new session, the results of the last session remain, which can not be said about users who started to fill the calendar without authorization.

2. Local own site: the application works without the use of the Internet, which is beneficial for users who do not have any means of online communication.

3. Optimization for conflicts: the calendar does not provide the same event headers, so in order to avoid disynchronization and conflicts, errors will be generated asking you to change the header and description

# How to use
- The application is launched through a local host on a special site that the developers have created. No Internet connection is required, so the calendar can be used even by a person who does not have Wi-Fi. However, the application itself can be launched in two ways: through the Docker virtual environment and through Ubuntu Linux. The first case is suitable for those who have a 64-bit operating system of any version, except for the home version: the latter version has no virtualization, which will prevent the calendar from launching. In the second case, it may be necessary to install Linux first, if the user does not have it (system or virtual machine). Let's demonstrate the principle of operation using the Linux case as an example.

The user should start the console and enter the following commands:

1. cd (name of the folder where the cloned project is located)
2. git pull
3. make build

After executing these three simple commands, the user will see a history of installed packages of the host in the console with a calendar and a link to the local website address to go to. The calendar itself and a few buttons to interact with it should appear in front of the user. However, in order to start performing all necessary operations, the user must register. This is necessary for a simple reason: all users who use the calendar create events for their own needs, and therefore registration will help to store a database for each of them, so that everything will be automatically saved during repeated sessions. Once registered, the user can perform all necessary actions.

The user can view data on current and upcoming events: he clicks on the appropriate buttons and the database is presented to him. He will also be able to add new events and edit them. To do this, he needs to select one of the days after the number that means the current day and enter the data, including the event name and description. It is crucial that the events have different titles, otherwise a conflict may occur. In this case, an error will be triggered, signaling the user to change the description slightly. The same procedure works when editing events. 