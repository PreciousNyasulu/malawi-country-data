select userId, json_arrayagg(areas) as completedSteps
from (
    select distinct userId,
           json_object('area',area,'games',(json_arrayagg(games))) as areas
    from (
        select distinct userId,area,
        json_object('id',game,'steps',(json_arrayagg(step))) as games
        from scoreTable
        where completed = 1 
        ) st1
    group by userId, area) st2
group by userId;