import{test}from'node:test';import assert from'node:assert/strict';import{matchesFilters,nearTitle,suggestions}from'./filters.ts';
const r={source:'tasks',id:'1',title:'Quarterly planning',snippet:'',href:'/tasks/',kind:'Tasks',updated:'2026-10-19T12:00:00Z',tags:['Work']};
test('supported filter fields and unknown metadata',()=>{assert.equal(matchesFilters(r,{after:'2026-10-01',tagFilter:'work'}),true);assert.equal(matchesFilters(r,{ownerFilter:'alice'}),false);assert.equal(matchesFilters({...r,updated:''},{after:'2026-10-01'}),false)});
test('single-edit spelling and authorised suggestions only',()=>{assert.equal(nearTitle('Planning launch','planing'),true);assert.equal(nearTitle('Private unrelated','planing'),false);assert.deepEqual(suggestions('plan',[r],[]),['Quarterly planning'])});
