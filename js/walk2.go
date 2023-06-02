package js

type IVisitor2 interface {
	Enter(n INode)
	Exit(n INode)
}

func Walk2(v IVisitor2, n INode) {
	visit(v, n, nil)
}

func visit(v IVisitor2, n INode, parent INode) {
	if n == nil {
		return
	}

	v.Enter(n)
	defer v.Exit(n)

	switch n := n.(type) {
	case *AST:
		visit(v, &n.BlockStmt, n)
	case *Var:
		return
	case *BlockStmt:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, n.List[i], n)
			}
		}
	case *EmptyStmt:
		return
	case *ExprStmt:
		visit(v, n.Value, n)
	case *IfStmt:
		visit(v, n.Body, n)
		visit(v, n.Else, n)
		visit(v, n.Cond, n)
	case *DoWhileStmt:
		visit(v, n.Body, n)
		visit(v, n.Cond, n)
	case *WhileStmt:
		visit(v, n.Body, n)
		visit(v, n.Cond, n)
	case *ForStmt:
		if n.Body != nil {
			visit(v, n.Body, n)
		}

		visit(v, n.Init, n)
		visit(v, n.Cond, n)
		visit(v, n.Post, n)
	case *ForInStmt:
		if n.Body != nil {
			visit(v, n.Body, n)
		}

		visit(v, n.Init, n)
		visit(v, n.Value, n)
	case *ForOfStmt:
		if n.Body != nil {
			visit(v, n.Body, n)
		}

		visit(v, n.Init, n)
		visit(v, n.Value, n)
	case *CaseClause:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, n.List[i], n)
			}
		}

		visit(v, n.Cond, n)
	case *SwitchStmt:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		visit(v, n.Init, n)
	case *BranchStmt:
		return
	case *ReturnStmt:
		visit(v, n.Value, n)
	case *WithStmt:
		visit(v, n.Body, n)
		visit(v, n.Cond, n)
	case *LabelledStmt:
		visit(v, n.Value, n)
	case *ThrowStmt:
		visit(v, n.Value, n)
	case *TryStmt:
		if n.Body != nil {
			visit(v, n.Body, n)
		}

		if n.Catch != nil {
			visit(v, n.Catch, n)
		}

		if n.Finally != nil {
			visit(v, n.Finally, n)
		}

		visit(v, n.Binding, n)
	case *DebuggerStmt:
		return
	case *Alias:
		return
	case *ImportStmt:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}
	case *ExportStmt:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		visit(v, n.Decl, n)
	case *DirectivePrologueStmt:
		return
	case *PropertyName:
		visit(v, &n.Literal, n)
		visit(v, n.Computed, n)
	case *BindingArray:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		visit(v, n.Rest, n)
	case *BindingObjectItem:
		if n.Key != nil {
			visit(v, n.Key, n)
		}

		visit(v, &n.Value, n)
	case *BindingObject:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		if n.Rest != nil {
			visit(v, n.Rest, n)
		}
	case *BindingElement:
		visit(v, n.Binding, n)
		visit(v, n.Default, n)
	case *VarDecl:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}
	case *Params:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		visit(v, n.Rest, n)
	case *FuncDecl:
		visit(v, &n.Body, n)
		visit(v, &n.Params, n)

		if n.Name != nil {
			visit(v, n.Name, n)
		}
	case *MethodDecl:
		visit(v, &n.Body, n)
		visit(v, &n.Params, n)
		visit(v, &n.Name, n)
	case *Field:
		visit(v, &n.Name, n)
		visit(v, n.Init, n)
	case *ClassDecl:
		if n.Name != nil {
			visit(v, n.Name, n)
		}

		visit(v, n.Extends, n)

		for _, item := range n.List {
			if item.StaticBlock != nil {
				visit(v, item.StaticBlock, n)
			} else if item.Method != nil {
				visit(v, item.Method, n)
			} else {
				visit(v, &item.Field, n)
			}
		}
	case *LiteralExpr:
		return
	case *Element:
		visit(v, n.Value, n)
	case *ArrayExpr:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}
	case *Property:
		if n.Name != nil {
			visit(v, n.Name, n)
		}

		visit(v, n.Value, n)
		visit(v, n.Init, n)
	case *ObjectExpr:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}
	case *TemplatePart:
		visit(v, n.Expr, n)
	case *TemplateExpr:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}

		visit(v, n.Tag, n)
	case *GroupExpr:
		visit(v, n.X, n)
	case *IndexExpr:
		visit(v, n.X, n)
		visit(v, n.Y, n)
	case *DotExpr:
		visit(v, n.X, n)
		visit(v, &n.Y, n)
	case *NewTargetExpr:
		return
	case *ImportMetaExpr:
		return
	case *Arg:
		visit(v, n.Value, n)
	case *Args:
		if n.List != nil {
			for i := 0; i < len(n.List); i++ {
				visit(v, &n.List[i], n)
			}
		}
	case *NewExpr:
		if n.Args != nil {
			visit(v, n.Args, n)
		}

		visit(v, n.X, n)
	case *CallExpr:
		visit(v, &n.Args, n)
		visit(v, n.X, n)
	case *UnaryExpr:
		visit(v, n.X, n)
	case *BinaryExpr:
		visit(v, n.X, n)
		visit(v, n.Y, n)
	case *CondExpr:
		visit(v, n.Cond, n)
		visit(v, n.X, n)
		visit(v, n.Y, n)
	case *YieldExpr:
		visit(v, n.X, n)
	case *ArrowFunc:
		visit(v, &n.Body, n)
		visit(v, &n.Params, n)
	case *CommaExpr:
		for _, item := range n.List {
			visit(v, item, n)
		}
	default:
		return
	}
}
