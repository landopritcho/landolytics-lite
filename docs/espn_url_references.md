# ESPN URL Reference Guide

## Games on a given date:
https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard
- dates (YYYYMMDD format)
- limit

## Game summary, including team and player box scores:
https://site.api.espn.com/apis/site/v2/sports/football/nfl/summary
- event (unique event ID provided by ESPN)

## Complete play-by-play:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/plays
- limit

## Drives and their plays:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/drives
- limit

## Current game situation:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/situation

## Detailed statistics for a team:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/competitors/{team_id}/statistics

## Quarter-by-quarter scores for a team:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/competitors/{team_id}/linescores

## Win probability by play:
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/{event_id}/competitions/{event_id}/probabilities
- limit

## CDN full game package:
https://cdn.espn.com/core/nfl/game
