## Metropolis in your Organization

> *Note*: In this chapter, 'developers' mean product developers, ie. Metropolis **Users**, not Metropolis **Developers**. Whenever you see **User**, **Operator** or **Developer**, think of Metropolis roles. However, whenever you see **developer**, think of product development teams acting as **Users**.

As outlined in [How to use this Handbook](introduction-01-how-to-use.md), Metropolis has at its core the concept of separate Users and Operators of a Metropolis cluster.

This split might, at first glance, seem antithetical to the spirit of 'DevOps'. However, this distinction **doesn't exist to take away operational tasks from software developers** (Users), but to let Metropolis scale to large organizations where developers cannot be expected to be responsible for operations from physical hardware (or a public cloud) up to their product. We believe product teams should be able to focus on the operational aspects specific to their product, and not have to deal with low-level fluff like cluster-level backups, monitoring and security.

This chapter aims to explain and argument the reasoning for such a split, and tie this into how Metropolis expects to be managed in different kinds of organizations.

### Platform Teams

Metropolis allows large organizations to build internal Platform Teams. These exist to bring a 'PaaS'-style experience to multiple internal product development teams. Metropolis neatly fits into this scenario by exposing only a standard Kubernetes API to these development teams (acting as Metropolis Users), while also exposing a powerful but proprietary API for the platform team (Metropolis Operators) that concerns only operational work. The two APIs are separate but do not overlap in functionality.

In the following example, the Platform Team are Metropolis Operators, while Product Teams A and B are Metropolis Users. The Platform Team runs two multi-tenant Metropolis clusters, both of which can be used by any Product Team for any purpose.

```
.----------------.      Manages     .--------------------------.
| Product Team A | ---------------> | Product A                |
'----------------'       (k8s)      '--------------------------'
                                         |            |
.----------------.      Manages     .--------------------------.
| Product Team B | ---------------> | Product B                |
'----------------'       (k8s)      '--------------------------'
                                         | |          | |
                                         | |          | | Runs on
                                         V V          V V
.----------------.      Manages     .-----------.  .-----------.
| Platform Team  | ---------------> | Cluster X |  | Cluster Y |
|                | -----------------|           |->|           |
'----------------' (Metropolis API) '-----------'  '-----------'
```

At large scales, Product Teams benefit from Metropolis by using a product that does not require them to be aware of implementation details below the Kubernetes API layer, and can focus on day-to-day operations of core products. They do not need to coordinate with other Product Teams on sharing the underlying resources, nor do they need to take care of managing or scaling the cluster. Platform Teams likewise benefit from Metropolis having been designed for use in a multi-tenant fashion where product teams can share a cluster safely.

### Smaller Organizations

While the above Platform Team system works great for larger organizations, smaller organizations usually do not benefit from having distinct teams of dozens of people responsible just for clusters and other organizational-wide resources.

In these cases, there is nothing which prevents the lone backend developer at a company from acting both as a Metropolis Operator and User and managing both Metropolis clusters and the actual workloads running on it:

```
.--------------.       Manages      .---------.
| Backend Team | -----------------> | Product |
|              |        (k8s)       |         |
|              | --.                '---------'
'--------------'   |                     | Runs on
                   |                     V
                   |   Manages      .---------.
                   '--------------> | Cluster |
                  (Metropolis API)  '---------'
```

As the organization grows, Metropolis will continue gently guiding (by way of Users/Operators role separation) workflows of the Backend team to not mix these two roles together. From the beginning, the Product can be deployed only using the Kubernetes API without needing to touch Metropolis-specific APIs. As new products and projects are developed, these can continue to use the existing Metropolis infrastructure without overhead of having each team manage their own production from the ground up.

### Organizational anti-patterns

Monogon believes that organizational issues cannot simply be fixed by applying technical solutions. Thus, Metropolis explicitly avoids supporting usecases that stem from heavy internal siloization of organizations, or the broken incentives of a syadmin-style platform team. We believe that Metropolis can be used as a catalyst to build better teams and workflows, but it is not by itself a fix for organizational problems.

We would like to refer you to the following sources for more information on these organizational patterns and anti-patterns.

1. [Infra teams: good, bad or none at all](https://rachelbythebay.com/w/2020/05/19/abc/), which describes the typical emerging ways organizations deal with infrastructure work. Metropolis leans heavily towards a “Company A” environment.
1. [The SRE Book](https://sre.google/sre-book/table-of-contents/), which describes Google's “implementation” of DevOps. While the processes described work best for extremely large companies, a significant amount of high-level observations and judgements can be pertinent to even the smallest organizations.
1. [The SRE Workbook](https://sre.google/workbook/table-of-contents/) chapter [“How SRE Relates to DevOps”](https://sre.google/workbook/how-sre-relates/), which describes an organizational approach to development and operation teams in which Metropolis works best.

