# Okanagna Information Security Group - K8s Security Presentation

## Presentation Resources 

* https://docs.microsoft.com/en-us/azure/aks/use-pod-security-policies
* https://www.openlogic.com/blog/granting-user-access-your-kubernetes-cluster
* https://docs.bitnami.com/kubernetes/how-to/secure-kubernetes-cluster-psp/
* https://kubernetes.io/docs/concepts/policy/pod-security-policy/#privileged

## Revealjs Info

### Markdown Formatting
A single markdown file is used for each section. 2 blank spaces indicates a new slide vertically, while 3 indicates a new slide horizontally. This can be changed in the index.html which is currently configured as: 

```
                <section data-markdown="content/00_intro/index.md"
                    data-separator="^\n\n\n"
                    data-separator-vertical="^\n\n"
                    data-separator-notes="^Note:"
                    data-charset="iso-8859-15">
                </section>
```

### Random Bacground Image

`for x in $(seq 0 6); do sleep 4 && wget -O slides/default/images/background-$x.jpg https://source.unsplash.com/1600x900/?security && convert slides/default/images/background-$x.jpg -fill 'rgb(24,54,105)' -colorize 60% slides/default/images/background-$x.jpg && composite -gravity southeast -geometry +50+75 slides/default/images/arctiq-logo-white.png slides/default/images/background-$x.jpg slides/default/images/background-$x.jpg; done`


### Local Usage
- Build a new container with updated content

```
docker build -t preso .
```

- Run the container locally specifying the desired content folder

```
docker run -it --rm -e PRESO_NAME=default -p 8000:8000 preso
```


### Resources and Credit
#### Resources
* https://github.com/hakimel/reveal.js/wiki/Plugins,-Tools-and-Hardware
* https://github.com/byteclubfr/uncloak (CSS Editor)

#### Credit
* https://github.com/nbrownuk