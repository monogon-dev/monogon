# How to use this Handbook

This handbook is the canonical documentation for Metropolis. It aims to document all aspects of Metropolis, from a quick demo, through production deployment to architecture internals.

> Note: **This section is critical to understand the Handbook structure** and must be read by anyone looking to use Metropolis. At the bottom of this page you will find information about which sections to read next, depending on how you want to use Metropolis.

## Who is this book for?

Throughout this book, we will keep using the following terminology regarding groups of people who interact with Metropolis. Note, that these names *do NOT imply that these groups are disjoint*. Instead, think of them as *different hats* people can wear when using Metropolis.

Metropolis does not enforce these roles explicitly, but is designed and engineered in a way that makes the most sense for this kind of organizational structure.

### Operators

Operators are responsible for **managing Metropolis clusters** - for example, bringing nodes into the cluster and decommissioning old nodes, monitoring resource usage and performing capacity planning, ensuring that the cluster is healthy, and responsible for handling cluster-level outages.

Operators are usually part of an organization-wide 'platform' team which acts as an internal 'as-a-Service' service, providing services to Users (like a workload scheduling system, a database service, a storage service...). They might manage some workloads running on Metropolis themselves, too, usually parts of the platform provided to Users (like running organization-wide database clusters on Metropolis).

Operators mostly act as system administrators, but are expected to also be able to use Metropolis APIs from a programming or scripting language of their choice to automate their work and make the most of Metropolis. Metropolis provides a set of management tools that allow management from a command line, but these are only thin wrappers around the underlying API which should be the main way to think about Metropolis.

### Users

Users **run workloads** on Metropolis clusters, via the Kubernetes API. They might know that a cluster runs Metropolis, but this is generally not something they need to worry about - instead, they should be aware of the abstractions which Kubernetes provides. Some limited amount of interaction with Metropolis-specific APIs might exist for purposes of authentication or accounting, but would also be usually hidden away by Operators as part of some organizational integration code.

Critically, however, Metropolis does not provide Users with some 'friendly' higher-level API or tooling which duplicates functionality of the Metropolis API used by Operators. Instead, the Kubernetes API is provided to both Users and Operators in the same fashion so that any internal tooling built on top can be shared between users and Operators.

In addition, Metropolis makes no attempt to hide that it itself and Kubernetes are distributed systems and applications running on top of clusters need to be engineered to handle such a scenario.

In most organizations, Users will be part of product teams, both developing and operating the organization's product or service.

### Developers

Developers **work on the Metropolis codebase**. Metropolis is an open source project, welcomes external contributions and attempts to have a fully open design process. However, any changes introduced must be carefully reviewed and tested - not only external contributions, but also contributions from Monogon employees.

Metropolis comes with high quality developer tooling to work on the codebase - all tests, including full cluster tests, can be run without any special software straight from a Monogon repository checkout.

People who wish to build Metropolis from source (for security, to reproduce official artifacts, or to apply internal organization patches) are also expected to fall into this category. In the future, purpose-specific documentation might be built for software packagers or people who wish to ensure Metropolis builds are reproducible, but that is not the case yet.

## Which sections should be read, and in what order?

If you just want to try out Metropolis, head over to **2. Demo Cluster** and come back here later.

If you're considering deploying Metropolis, you must read **1. Metropolis in your Organization**. It lays some ground concepts of how Metropolis will fit in your organization, what it's good at, and what it's not. It's aimed towards future **Operators** that wish to better understand the relationship between them and Users, but should also be read by organization management teams that will oversee future Operator and User teams/roles.

**Operators** must read the following chapters:

  - **3. Cluster Architecture**, which describes how Metropolis is designed. The information contained therein is crucial to properly plan, deploy and manage a Metropolis cluster. Individual sections of the chapter will be marked if some information is optional in some kinds of deployments, and these parts might be skipped and read as needed later.
  - **4. Production Deployment**, which describes the standard procedures used to manage Metropolis clusters, troubleshooting procedures.
 
**Users** do not need to read any Metropolis-specific documentation to use Metropolis, and instead should rely on information provided by cluster Operators and the upstream Kubernetes documentation. However, we encourage users to skim through **3. Cluster Architecture** if they are interested in knowing more about the internal Metropolis architecture.

**Developers** are generally expected to start out as **Operators** and thus have read all relevant documentation for Operators already. In addition to that, they are provided with information on how to develop Metropolis in **5. Developing Metropolis**, which gives an introduction to the Metropolis codebase and how to get started writing code.
