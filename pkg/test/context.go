// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/ramendr/ramen/e2e/types"
)

// Context implements types.TestContext interface.
type Context struct {
	ctx      types.Context
	workload types.Workload
	deployer types.Deployer
	name     string
	logger   *zap.SugaredLogger
}

var _ types.TestContext = &Context{}

func newContext(ctx types.Context, workload types.Workload, deployer types.Deployer) *Context {
	name := fmt.Sprintf("%s-%s", deployer.GetName(), workload.GetName())
	return &Context{
		ctx:      ctx,
		workload: workload,
		deployer: deployer,
		name:     name,
		logger:   ctx.Logger().Named(name),
	}
}

func (c *Context) Deployer() types.Deployer {
	return c.deployer
}

func (c *Context) Workload() types.Workload {
	return c.workload
}

func (c *Context) Name() string {
	return c.name
}

func (c *Context) ManagementNamespace() string {
	if ns := c.deployer.GetNamespace(c); ns != "" {
		return ns
	}
	return c.AppNamespace()
}

func (c *Context) AppNamespace() string {
	return namespacePrefix + c.name
}

func (c *Context) Logger() *zap.SugaredLogger {
	return c.logger
}

func (c *Context) Env() *types.Env {
	return c.ctx.Env()
}

func (c *Context) Config() *types.Config {
	return c.ctx.Config()
}

func (c *Context) Context() context.Context {
	return c.ctx.Context()
}
