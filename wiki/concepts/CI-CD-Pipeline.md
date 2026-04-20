# CI/CD Pipeline

**Status**: Core Standard
**Type**: Engineering Pattern
**Parent**: [[DevOps]]
**Source**: [[redhat-cicd-pipeline]]

## 📝 Definition
A CI/CD pipeline is a series of automated steps that developers follow to deliver new versions of software. It focuses on improving software delivery throughout the lifecycle via automation.

## 🏗️ The Pipeline Stages

### 1. Continuous Integration (CI)
- **Code**: Developer pushes code to a repo.
- **Build**: Automated compilation of the source.
- **Test**: Automated unit and integration tests (see [[sdd-verify]]).

### 2. Continuous Delivery (CD)
- **Release**: Automatically packaging the application (e.g., Docker image).
- **Staging**: Deploying to a test environment.

### 3. Continuous Deployment (CD)
- **Production**: Automated deployment to the live environment if all tests pass.

## 🛡️ CI/CD Security (DevSecOps)
- **Scanning**: Automated vulnerability scans (see [[Trivy]]).
- **Secrets**: Detecting hardcoded credentials (see [[Security-Guardrails]]).
- **Policies**: Enforcing compliance checks before merge.

---
**Last Updated**: 2026-04-19
