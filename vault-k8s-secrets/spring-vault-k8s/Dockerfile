FROM maven:3.5-jdk-8-alpine
ENV MVN_PROFILE=placeholder
ENV ROLE_ID=69907c2e-b6fe-0daf-4ea5-dc8b9cdfa85c
ENV SECRET_ID=placeholder
COPY app/ .
CMD ./mvnw spring-boot:run -Dspring-boot.run.profiles=${MVN_PROFILE} ; tail -f /dev/null
#ENTRYPOINT ["java","-jar","targer.gs-vault-config-0.1.0.jar"]
