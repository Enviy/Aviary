# aviary
<pre>
                                                            
                                                            
       .----.     .----..--.                                
        \    \   /    / |__|                .-.          .- 
         '   '. /'   /  .--.          .-,.--.\ \        / / 
    __   |    |'    /   |  |    __    |  .-. |\ \      / /  
 .:--.'. |    ||    |   |  | .:--.'.  | |  | | \ \    / /   
/ |   \ |'.   `'   .'   |  |/ |   \ | | |  | |  \ \  / /    
`" __ | | \        /    |  |`" __ | | | |  '-    \ `  /     
 .'.''| |  \      /     |__| .'.''| | | |         \  /      
/ /   | |_  '----'          / /   | |_| |         / /       
\ \._,\ '/                  \ \._,\ '/|_|     |`-' /        
 `--'  `"                    `--'  `"          '..'         
</pre>

Azure Function App for interacting with Twitter.

Note:
- This project assumes you have an Azure Function App.
- This project assumes you have a registered Twitter app.

Use:
- Review and update parameters in config.yaml.
- Review and update function parameters in Make file.
- Build using <pre>make build</pre>
- Deploy using <pre>make deploy</pre>

Current:
- Tag users and tweet a random new fox image on cron schedule.

In progress:
- Tweet Lovecraft-like blurbs. (RNN text prediction)

