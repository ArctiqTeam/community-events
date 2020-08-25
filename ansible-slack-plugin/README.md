# POC custom Ansible callback plugin 
This callback plugin will update a slack channel on each playbook run.
This is just a quick and dirty poc, if you are looking for a more solid foundation, please use https://docs.ansible.com/ansible/latest/plugins/callback/slack.html

## Requirements

- pip3 install ansible slack_webhook
- vagrant
- admin access to a Slack instance


## Install
- Create a channel
- Add a Webhook integration https://<your-slack-instance>.slack.com/apps/manage
- Select Custom Integration then Incoming WebHooks
- Click Add to Slack
- Choose the channel you want to post to
- Click "Add Incoming Webhooks Integration"
- Potentially Change the name and the Icon
- Initialize a $SLACK_WEBHOOK_URL environment variable with the Slack Hook Url

## Run
source .env
vagrant up
