use std::collections::HashMap;

use regex::Regex;

#[derive(PartialEq, Eq, Hash, Default, Debug)]
pub struct RuleItem {
    pub chemical: usize,
    pub quantity: usize,
}

#[derive(PartialEq, Eq, Hash, Default, Debug)]
pub struct ProductionRule {
    pub ingredients: Vec<RuleItem>,
    pub output: RuleItem,
}

#[derive(Debug)]
struct ProductionRuleParser<'a> {
    chemicals_mapping: HashMap<&'a str, usize>,
    item_regex: Regex,
}

pub struct ParsingResult {
    pub rules: Vec<ProductionRule>,
    pub fuel_index: usize,
    pub ore_index: usize,
}

pub fn parse_production_rules(lines: &[&str]) -> Result<ParsingResult, String> {
    let mut parser = ProductionRuleParser::new();

    Ok(ParsingResult {
        rules: lines
            .into_iter()
            .map(|l| parser.parse_rule(l))
            .collect::<Result<Vec<_>, _>>()?,
        fuel_index: parser.chemicals_mapping["FUEL"],
        ore_index: parser.chemicals_mapping["ORE"],
    })
}

impl<'a> ProductionRuleParser<'a> {
    fn new() -> Self {
        Self {
            chemicals_mapping: HashMap::new(),
            item_regex: Regex::new(r"\d+ \w+").expect("cannot compile item regexp"),
        }
    }

    fn parse_rule(&mut self, s: &'a str) -> Result<ProductionRule, String> {
        let mut matches = self
            .item_regex
            .find_iter(s)
            .collect::<Vec<_>>()
            .into_iter()
            .map(|m| self.parse_rule_item(m.as_str()))
            .collect::<Result<Vec<_>, _>>()?;

        Ok(ProductionRule {
            output: matches.remove(matches.len() - 1),
            ingredients: matches,
        })
    }

    fn parse_rule_item(&mut self, s: &'a str) -> Result<RuleItem, String> {
        let mut parts = s.split(' ');
        let quantity = parts
            .next()
            .ok_or("missing quantity")?
            .parse::<usize>()
            .map_err(|e| format!("cannot parse quantity: {}", e))?;

        let name = parts.next().ok_or("missing name")?;
        let next_chemical_index = self.chemicals_mapping.len();
        let chemical_index = self
            .chemicals_mapping
            .entry(name)
            .or_insert(next_chemical_index);

        Ok(RuleItem {
            chemical: *chemical_index,
            quantity,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn rule_parser_parses_correctly_single_ingredient() {
        let mut parser = ProductionRuleParser::new();
        let res = parser
            .parse_rule(
                "1 ORE => 1 B", // 7 A, 1 B => 1 C",
            )
            .expect("cannot parse rule");

        assert_eq!(
            res.ingredients,
            vec![RuleItem {
                chemical: 0,
                quantity: 1
            }]
        );
        assert_eq!(
            res.output,
            RuleItem {
                chemical: 1,
                quantity: 1
            }
        );
    }

    #[test]
    fn rule_parser_parses_correctly_multiple_ingredients() {
        let mut parser = ProductionRuleParser::new();
        let res = parser
            .parse_rule("1 ORE, 1 B, 2 C => 1 B")
            .expect("cannot parse rule");

        assert_eq!(
            res.ingredients,
            vec![
                RuleItem {
                    chemical: 0,
                    quantity: 1
                },
                RuleItem {
                    chemical: 1,
                    quantity: 1
                },
                RuleItem {
                    chemical: 2,
                    quantity: 2
                }
            ]
        );
        assert_eq!(
            res.output,
            RuleItem {
                chemical: 1,
                quantity: 1
            }
        );
    }
}
