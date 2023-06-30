# ReST (Representational State Transfer)

CRUD => Create Read Update Destroy

HTTP Methods => {GET, PUT, POST, OPTIONS, PATCH, ...}

## Employee Management Server (JSON)

```
CRUD         Action             HTTP Method                   URI                   Request Body                     Response Body
-----------------------------------------------------------------------------------------------------------------------------------
Read         Index                GET                     /employees                    -                               [{...}, ...] -- Done
Read         Show                 GET                     /employees/{id}               -                               {...}        -- Done
Create       Create               POST                    /employees                  {...}                             {id: , ...}  -- Done
Update       Update               PUT                     /employees/{id}             {...}                             {...}
Update       Update               PATCH                   /employees/{id}             {selected attrs}                  {...}
Destroy      Destroy              DELETE                  /employess/{id}               -                                - / {...}
```
