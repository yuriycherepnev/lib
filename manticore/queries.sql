SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('h18', name, {normalize=1})*100) as levenshtein
FROM tyres_search WHERE MATCH('(h18 | *h18*)') 
AND ANY(levenshtein)<95 ORDER BY levenshtein ASC, weight DESC LIMIT 100;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('nokian', name, {normalize=1})*100) as levenshtein
FROM tyres_search 
WHERE MATCH('(*nokian*)') 
AND ANY(id_search_category)=1 
AND ANY(levenshtein)<95 
ORDER BY levenshtein ASC, weight DESC LIMIT 100;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('pireli', name, {normalize=1})*100) as levenshtein
FROM tyres_search 
WHERE MATCH('(pireli)') 
ORDER BY levenshtein ASC, weight DESC LIMIT 100;


SELECT *
  FROM (
   SELECT tyres_search.*, num:=if(man=id_search_category,num+1,1) as NUM, man:=id_search_category
     FROM tyres_search, (select man:='',num:=0) A

     WHERE MATCH('(pireli)') 

    ORDER BY id_search_category
  ) A
WHERE NUM<=2;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('belshinak', name, {normalize=1})*100) as levenshtein
FROM tyres_search 
WHERE MATCH('(belsh | *belsh*)') 
ORDER BY levenshtein ASC, weight DESC LIMIT 5;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('y e', name, {normalize=1})*100) as levenshtein
FROM tyres_search WHERE MATCH('(ь | *ь*) & (е | *е*)') AND ANY(id_search_category)=1 AND ANY(levenshtein)<95 
ORDER BY levenshtein ASC, weight DESC LIMIT 0;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('235 16', name, {normalize=1})*100) as levenshtein
FROM tyres_search 
WHERE MATCH('(*235*) & (16 | r16)') 
AND ANY(id_search_category)=1 
AND ANY(levenshtein)<95 
ORDER BY weight DESC, levenshtein ASC LIMIT 100;


SELECT *, weight()*custom_weight as weight, ceil(LEVENSHTEIN('zshkudsh', name, {normalize=1})*100) as levenshtein
FROM tyres_search WHERE MATCH('(*Зшкудш*)') 
ORDER BY weight DESC, levenshtein ASC LIMIT 5;


SELECT name, custom_weight, weight()*custom_weight as weight
FROM tyres_search 
WHERE MATCH('(195 60)') 
ORDER BY weight DESC LIMIT 5;