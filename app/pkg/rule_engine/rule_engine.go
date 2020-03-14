package rule_engine

import (
	grule "github.com/hyperjumptech/grule-rule-engine"
)

func InitRuleEnginee() (*grule.engine.NewGruleEngine(), error) {
	engine := grule.engine.NewGruleEngine()
	err := engine.Execute(dataCtx, knowledgeBase, workingMemory)
	if err == nil {
		return engine, nil
	}
	return nil, err
}

func loadRules() {
	// https://github.com/hyperjumptech/grule-rule-engine/blob/master/docs/Tutorial_en.md
	// for file in rule_dirs {
	// 	fileRes := grule.pkg.NewFileResource("/path/to/rules.grl")
	// 	err := ruleBuilder.BuildRuleFromResource(fileRes)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// bundle
	bundle := grule.pkg.NewFileResourceBundle("/path/to/grls", "/path/to/grls/**/*.grl")
	resources := bundle.MustLoad()
	for _, res := range resources {
		err := ruleBuilder.BuildRuleFromResource(res)
		if err != nil {
			panic(err)
		}
	}
}
