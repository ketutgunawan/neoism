// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package neo4j

import (
	"testing"
	// "fmt"
)

// 18.3.1. Send queries with parameters
func TestCypherSendQueryWithParameters(t *testing.T) {
	// Create
	// idx0, _ := db.Nodes.Indexes.CreateWithConf("name_auto_index", "fulltext", "lucene")
	n0, _ := db.Nodes.Create(Properties{"name": "I"})
	n1, _ := db.Nodes.Create(Properties{"name": "you"})
	r0, _ := n0.Relate("know", n1.Id(), nil)
	// Deferred Cleanup
	defer r0.Delete()
	defer n0.Delete()
	defer n1.Delete()
	// defer idx0.Delete()
	// Query
	query := "START x = node:node_auto_index(name={startName}) MATCH path = (x-[r]-friend) WHERE friend.name = {name} RETURN TYPE(r)"
	// query := fmt.Sprintf("START n=node(%v) RETURN n", n0.Id())
	params := map[string]string{
		// "startName": "I",
		// "name":      "you",
	}
	result, err := db.Cypher(query, params)
	if err != nil {
		t.Error(err)
	}
	// Check result
	logPretty(result)
}
