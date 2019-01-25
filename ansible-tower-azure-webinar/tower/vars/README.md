### Extra Vars

All the vars with no entry in the extra_vars.yml file are essentially secrets that shouldn't be public, so you should have 
your secrets put in place of those empty variables when using this role.

There is a `.gitignore` file in this role that omits a secrets.yml file from being added to this repo, to avoid letting those go public.
