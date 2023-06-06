Schema/Version compatibility
===

Live migration
---

BMaaS supports live migrating schemas. On startup, every component using the BMaaS
will attempt to migrate the database up to the newest version of the schema it
was built with.

Components are implemented to support a range of schemas, and operators should
sequence upgrades in the following way:

1. Make sure that all components are at the newest possible CL, but not so new
   that they ship a newer version of the schema than is currently running.
2. Upgrade components in a rolling fashion to a CL version that ships the newest
   possible schema version which is still compatible with the previous CL
   versions of the components.
3. Repeat from point 1 until the newest wanted CL version is running.

| ID | Schema range  | CL range | Notes                        |
|----|---------------|----------|------------------------------|
| 0  | < 1672749980  | >= 0     | Initial production schema.   |
| 1  | >= 1672768890 | >= 1565  | Exponential backoff support. |

For example, if the cluster is at version 1200, it should first be upgraded to 
< 1565 (to reach row 0), then to anything higher than 1565 (to reach row 1).

Offline migration
---

For simple deployments, an offline migration is easiest. To perform an offline migration:

1. Turn down all BMaaS components that communicate with the BMDB.
2. Upgrade all components to the newer version (either newest or otherwise
   wanted, but all components have to be at the same CL version).
3. Turn up a single component of BMaaS torn down in 1., making sure the database is migrated.
4. Turn up the rest of the components.

This allows migrating across many incompatible schema migrations, but requires downtime.