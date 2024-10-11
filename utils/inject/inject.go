package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/facebookgo/inject"
)

type HomePlanetRenderApp struct {
	// 下面的标签向注入库表明这些字段符合注入条件。它们不指定任何选项，并将导致为每个 API 创建一个单例实例。
	NameAPI   *NameAPI   `inject:""`
	PlanetAPI *PlanetAPI `inject:""`
}

func (a *HomePlanetRenderApp) Render(id uint64) string {
	return fmt.Sprintf(
		"%s is from the planet %s.",
		a.NameAPI.Name(id),
		a.PlanetAPI.Planet(id),
	)
}

type NameAPI struct {
	// 在 PlanetAPI 中，我们将标签添加到接口值中。
	// 该值无法自动创建（根据定义），因此必须显式提供给图表。
	HTTPTransport http.RoundTripper `inject:""`
}

func (n *NameAPI) Name(id uint64) string {
	// 我们将使用 f.HTTPTransport 并获取名称
	return "Spock"
}

type PlanetAPI struct {
	HTTPTransport http.RoundTripper `inject:""`
}

func (p *PlanetAPI) Planet(id uint64) string {
	return "Vulcan"
}

func main() {
	var g inject.Graph

	// 我们为图表提供两个“种子”对象，一个是我们希望填充的空 HomePlanetRenderApp 实例，
	// 第二个是我们的 DefaultTransport 以满足我们的 HTTPTransport 依赖关系。
	// 我们必须提供 DefaultTransport，因为依赖项是根据 http.RoundTripper 接口定义的，
	// 并且由于它是一个接口，因此库无法为其创建实例。相反，它将使用给定的 DefaultTransport 来满足依赖关系，因为它实现了该接口：
	var a HomePlanetRenderApp
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: http.DefaultTransport},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 这里的 Populate 调用正在创建 NameAPI 和 PlanetAPI 的实例，
	// 并将两者上的 HTTPTransport 设置为上面提供的 http.DefaultTransport：
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 也可以使用这个短的API，它相当于上面2步
	//inject.Populate(&a, http.DefaultTransport)

	fmt.Println(a.Render(42))
}
