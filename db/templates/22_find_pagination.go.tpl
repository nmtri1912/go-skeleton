{{- $alias := .Aliases.Table .Table.Name}}

// Paginate returns filtered {{$alias.UpSingular}} records from the query within limit count.
func GetAll{{$alias.UpSingular}}s(ctx context.Context, exec boil.ContextExecutor, params FilterParam, selectCols ...string) ({{$alias.UpSingular}}Slice, int64, error) {
	{{$alias.DownSingular}}Obj := []*{{$alias.UpSingular}}{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	filterQuery := []string{}
	if len(params.Filters) > 0 {
		for _, filter := range params.Filters {
			filterQuery = append(filterQuery, filter.ToQuery())
		}
	}
	where := "1=1"
	if len(filterQuery) > 0 {
		where = strings.Join(filterQuery, "AND")
	}
	if params.Limit == 0 { // Set default limit
		params.Limit = 10
	}

	query := fmt.Sprintf(
		"select %s from {{.Table.Name | .SchemaTable}} where %s and \"deleted_at\" is null LIMIT $1 OFFSET $2", sel, where,
	)

	q := queries.Raw(query, params.Limit, params.Offset)

	err := q.Bind(ctx, exec, &{{$alias.DownSingular}}Obj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, 0, sql.ErrNoRows
		}
		return nil, 0, errors.Wrap(err, "models: unable to select from {{$alias.DownPlural}}")
	}

    // Count total
	var totalCount int64
	countQuery := queries.Raw(
		fmt.Sprintf(
			"select COUNT(*) from {{.Table.Name | .SchemaTable}} where %s and \"deleted_at\" is null", where,
		),
	)
	err = countQuery.QueryRowContext(ctx, exec).Scan(&totalCount)
	if err != nil {
		return nil,totalCount, errors.Wrap(err, "models: failed to count products rows")
	}

	if len({{$alias.DownSingular}}AfterSelectHooks) != 0 {
		for _, obj := range {{$alias.DownSingular}}Obj {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return {{$alias.DownSingular}}Obj, totalCount, err
			}
		}
	}

	return {{$alias.DownSingular}}Obj, totalCount, nil
}