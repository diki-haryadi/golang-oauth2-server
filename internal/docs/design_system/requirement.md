# Functional Requirement
# Functional Requirement
- as an user i can login seamless
- as an user i can see my last interview session with details
- as an user i can start interview assistant
- as an user i can speech and got AI response from gpt-model
- as an user i can see my speech to text base and question answer
- as an user i can select translation question and answer
- as an user i can enable disable autoscrolling AI response
- as an user i can direct call to AI with special button
- as an user i can see candidate score based question and answer
- as an user i can input every what hit to AI, title, custom intruction to AI
- as an system can record session and accumulate and decrease package
- as an user i can select package
- as an user i can pay gateway with any payment gateway type
- as an user i can select language listening and language answering
- as an user i can see my profile, account plan, troubleshoot, and usefull link

- more benefit
- as an user i can see suggest answer based question with AI responnse
- as an user i can see translate on question and answer on interview transcript
- as an user i can submit question to AI response
- as an user i can open browser in app

# Microservices Architecture

Based on the functional requirements you've outlined, we can design a microservices architecture that encompasses various services to handle the different functionalities. Here’s a breakdown of potential services:

## Microservices

1. **Authentication Service**
    - Handles user login, registration, and authentication.
    - Manages user sessions and tokens.

2. **User Profile Service**
    - Manages user profiles, account plans, and settings.
    - Provides endpoints to view and update user information.

3. **Interview Management Service**
    - Manages interview sessions, including starting sessions and retrieving past sessions.
    - Records session details and accumulates usage data for billing.

4. **AI Interaction Service**
    - Handles speech-to-text conversion and AI response generation.
    - Manages direct calls to the AI model and translates questions/answers.

5. **Translation Service**
    - Provides translation functionalities for questions and answers.
    - Integrates with the AI Interaction Service for seamless translation.

6. **Payment Gateway Service**
    - Manages package selection and payment processing.
    - Integrates with various payment gateways for user transactions.

7. **Session Recording Service**
    - Records audio/video sessions and manages storage.
    - Tracks usage for billing and package management.

8. **Recommendation Service**
    - Suggests answers based on user questions and AI responses.
    - Provides insights and analytics on user interactions.

9. **Browser Integration Service**
    - Allows users to open a browser within the app.
    - Manages browser sessions and integrates with the main application.

10. **Notification Service**
    - Sends notifications to users regarding updates, session reminders, and other alerts.

## Service Interactions

- **User Interaction Flow**:
    - Users authenticate via the Authentication Service.
    - They access their profile through the User Profile Service.
    - Users start an interview via the Interview Management Service.
    - During the interview, the AI Interaction Service handles speech input and provides AI responses.
    - The Translation Service handles any language translations needed.
    - Users can submit questions directly to the AI through the AI Interaction Service.
    - Session details are recorded and managed by the Session Recording Service.
    - Users can select and pay for packages via the Payment Gateway Service.

## Benefits of Microservices Approach

- **Scalability**: Each service can be scaled independently based on demand.
- **Flexibility**: Different technologies can be used for different services based on requirements.
- **Maintainability**: Smaller codebases make it easier to manage and update services.
- **Resilience**: Failure in one service does not affect the entire system, improving overall stability.

## Summary

This microservices architecture effectively addresses the functional requirements while providing a robust and scalable solution. Each service focuses on a specific aspect of the application, enabling better management and development agility.

# Design
https://excalidraw.com/#json=syK_q4fWGtNZ3QamqGc0w,yDH5e2-6drYg23Mzww_GMw

