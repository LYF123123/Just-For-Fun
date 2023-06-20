// Coding Best Practices: code must be self-documenting.
// so that even a brief look at it may tell you what it does.

// Me: Uhm... Like this?

p=[]    ,(f=    (n,   a     )   =>(!
a   &   n   >    2    &&  f(    n
-   1   ,   0    )    , p [ a   ]
?n%p    [a]?     f    (  n  ,   a+1
)       :   0    :    p     [   a
]       =   n   ,p)   )     (   1e2)

console.log(...p)