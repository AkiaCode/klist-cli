use comfy_table::Table;
use dialoguer::{theme::ColorfulTheme, Input};

pub fn search() {
    let input: String = Input::with_theme(&ColorfulTheme::default())
        .with_prompt("봇 아이디를 입력하세요")
        .validate_with(|input: &String| {
            if input.parse::<i128>().is_err() {
                Err("This is a String")
            } else {
                Ok(())
            }
        })
        .interact_text()
        .unwrap();
    let bot = koreanbots::blocking::Client::new("").get_bot(&input);
    if bot.code != 200 {
        return;
    }
    let mut table = Table::new();
    table
        .set_header(vec!["", "봇"])
        .add_row(vec!["이름", &bot.data.name])
        .add_row(vec!["플래그", &bot.data.flags.to_string()])
        .add_row(vec!["태그", &bot.data.tag])
        .add_row(vec!["라이브러리", &bot.data.lib])
        .add_row(vec!["접두사", &bot.data.prefix])
        .add_row(vec!["투표 수", &bot.data.votes.to_string()])
        .add_row(vec!["서버 수", &bot.data.servers.to_string()])
        .add_row(vec!["배너", &bot.data.banner]);

    println!("{}", table);
}
