from slack_webhook import Slack
import time
import datetime
import os
import sys
import socket
from ansible.plugins.callback import CallbackBase

webhook_url = os.environ.get("SLACK_WEBHOOK_URL")

class CallbackModule(CallbackBase):
    user = os.getenv('USER')
    hostname = socket.gethostname()
    playbook_name = os.path.basename(sys.argv[-1])
    start_time = time.time()
    start_datetime = datetime.datetime.now()

    def playbook_on_stats(self, stats):
        if stats.failures:
            message = "Failure"
            color = 'danger'
        else:
            message = "Success"
            color = 'good'

        execution_time = time.time()-self.start_time

        slack = Slack(url=webhook_url)
        message = "{}@{}\nPlaybook: {}".format(self.user, self.hostname, self.playbook_name)

        fields = [{
                    "title": "Start time",
                    "value": "{}".format(self.start_datetime),
                    "short": True
                },{
                    "title": "Execution time",
                    "value": "{:10.2f}s".format(execution_time),
                    "short": True
                }]
        
        for cat in ("ok","changed","unreachable","failures","skipped","rescued","ignored"):
            results = getattr(stats, cat, {})
            found = False
            if results:
                for key in getattr(stats, cat, {}).keys():
                    found = True
                    fields.append(
                        {
                            "title": cat.capitalize(),
                            "value": "{}".format(getattr(stats, cat, {}).get(key)),
                            "short": True
                        }
                    )
            if not found:
                fields.append(
                    {
                        "title": cat.capitalize(),
                        "value": "0",
                        "short": True
                    }
                )

        self._display.warning(stats.__dict__)
        payload = [{
            "fallback": "Ansible Play Recap for {}".format(self.playbook_name),
            "color": str(color),
            "title": "Ansible Play Recap : {}".format(self.playbook_name),
            "text": str(message),
            "fields": fields
        }]

        slack.post(attachments=payload)