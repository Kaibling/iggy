-- name: SaveWorkflowChildren :exec
INSERT INTO workflows_children (
  workflow_id, children_id
) VALUES (
  $1, $2
);

-- name: DeleteWorkflowChildren :exec
DELETE from workflows_children
where workflows_children.workflow_id = $1;