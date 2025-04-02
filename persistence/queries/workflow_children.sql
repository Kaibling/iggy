-- name: SaveWorkflowChildren :exec
INSERT INTO workflows_children (
  workflow_id, children_id, step_order
) VALUES (
  $1, $2, $3
);

-- name: DeleteWorkflowChildren :exec
DELETE from workflows_children
where workflows_children.workflow_id = $1;