# boot.fetch

Boot.fetch was inspired by a friend who wished there was someone monitoring his progress, and told him to shut up and code whenever he was on the boot.dev discord without progressing. Some other users appreciated the aggressive approach of our inside joke and I expanded it to act as an admin monitoring interface for an entire group of people being held accountable, and to gather more detailed stats than boot.dev offers.

To add custom ascii art to a user you can add to switch statement in Users.go. Users are added via the api/v1/Users endpoint. .env specifies local database connection options and web server port. The spacing on the graph may begin to overlap names or place it over needed info on the lins if there are a lot of users, and graph.py has instructions on how to create custom off-sets per user for a presentable final product. The file writer will generate detailed logs of every event possible in the case of weird
