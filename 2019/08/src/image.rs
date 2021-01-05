use std::{convert::TryFrom, error::Error, fmt::Display};

#[derive(Debug, PartialEq)]
pub enum ImageParseError {
    MismatchedSize { layer_size: usize, data_size: usize },
    NonDigitCharacter { c: char },
    UnknownError(String),
}

impl Display for ImageParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::MismatchedSize {
                data_size,
                layer_size,
            } => write!(
                f,
                "mismatched image size (data size {} with layer size {}, reminder: {}",
                data_size,
                layer_size,
                data_size % layer_size
            ),
            Self::NonDigitCharacter { c } => {
                write!(f, "non digit character {} cannot be parsed", c)
            }
            Self::UnknownError(reason) => write!(f, "unknown error: {}", reason),
        }
    }
}

impl Error for ImageParseError {}

enum Color {
    Black = 0,
    White = 1,
    Transparent = 2,
}

impl ToString for Color {
    fn to_string(&self) -> String {
        match self {
            Self::Transparent => " ",
            Self::Black => " ",
            Self::White => "X",
        }
        .to_owned()
    }
}

impl TryFrom<u32> for Color {
    type Error = String;

    fn try_from(value: u32) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Color::Black),
            1 => Ok(Color::White),
            2 => Ok(Color::Transparent),
            _ => Err(format!("invalid color {}", value).to_owned()),
        }
    }
}

#[derive(Debug)]
pub struct Image {
    width: usize,
    height: usize,
    pub layers: Vec<Vec<u32>>,
}

impl Image {
    pub fn parse(data: &str, width: usize, height: usize) -> Result<Self, ImageParseError> {
        let layer_size = width * height;

        if data.len() % layer_size != 0 {
            return Err(ImageParseError::MismatchedSize {
                data_size: data.len(),
                layer_size,
            });
        }

        let layers = data.len() / layer_size;

        let mut image = Self {
            width,
            height,
            layers: vec![],
        };

        for i in 0..layers {
            let layer_start = layer_size * i;
            let layer: Vec<u32> = data
                .get(layer_start..(layer_start + layer_size))
                .ok_or(ImageParseError::UnknownError(
                    format!(
                        "could not get layer of length {} starting at {}",
                        layer_size, layer_start
                    )
                    .to_owned(),
                ))
                .and_then(|layer| {
                    layer
                        .chars()
                        .map(|c| {
                            c.to_digit(10)
                                .ok_or(ImageParseError::NonDigitCharacter { c })
                        })
                        .collect()
                })?;

            image.layers.push(layer);
        }

        return Ok(image);
    }

    pub fn get_composed_layer(&self) -> Option<Layer> {
        let mut final_layer = Layer {
            width: self.width,
            data: self.layers.get(0)?.clone(),
        };

        for layer in self.layers[1..].iter() {
            layer.iter().enumerate().for_each(|(i, d)| {
                if final_layer.data[i] == (Color::Transparent as u32)
                    && *d != (Color::Transparent as u32)
                {
                    final_layer.data[i] = *d;
                }
            })
        }

        Some(final_layer)
    }
}

pub struct Layer {
    width: usize,
    data: Vec<u32>,
}

impl ToString for Layer {
    fn to_string(&self) -> String {
        let lines: Vec<String> = self
            .data
            .chunks_exact(self.width)
            .map(|line| {
                line.iter()
                    .map(|c| match Color::try_from(*c) {
                        Ok(color) => color.to_string(),
                        Err(_) => "E".to_owned(),
                    })
                    .collect::<Vec<_>>()
                    .join("")
            })
            .collect();

        lines.join("\n")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parses_image_correctly() {
        let data = "123456789012";
        let image = Image::parse(data, 3, 2).expect("Image parsing error");

        assert_eq!(image.width, 3, "invalid width");
        assert_eq!(image.height, 2, "invalid height");
        assert_eq!(image.layers.len(), 2, "invalid number of layers");
        assert_eq!(
            image.layers[0],
            vec![1, 2, 3, 4, 5, 6],
            "invalid first layer"
        );
        assert_eq!(
            image.layers[1],
            vec![7, 8, 9, 0, 1, 2],
            "invalid second layer"
        );
    }

    #[test]
    fn detects_invalid_image_characters() {
        let data = "123456789ab2";
        let error = Image::parse(data, 3, 2)
            .expect_err("Image parsing succeeded when an error was expected");

        assert_eq!(
            error,
            ImageParseError::NonDigitCharacter { c: 'a' },
            "invalid error returned"
        );
    }

    #[test]
    fn detects_invalid_image_length() {
        let data = "1234567";
        let error = Image::parse(data, 3, 2)
            .expect_err("Image parsing succeeded when an error was expected");

        assert_eq!(
            error,
            ImageParseError::MismatchedSize {
                data_size: 7,
                layer_size: 6
            },
            "invalid error returned"
        );
    }

    #[test]
    fn composes_layers() {
        let data = "0222112222120000";
        let image = Image::parse(data, 2, 2).expect("Image parsing error");
        let layer = image
            .get_composed_layer()
            .expect("Could not compose layers");

        assert_eq!(layer.width, 2, "invalid width");
        assert_eq!(layer.data, vec![0, 1, 1, 0], "invalid composed layer");
    }

    #[test]
    fn composed_layer_stringifies() {
        let data = "0222112222120000";
        let image = Image::parse(data, 2, 2).expect("Image parsing error");
        let layer = image
            .get_composed_layer()
            .expect("Could not compose layers");

        assert_eq!(
            layer.to_string(),
            " X
X ",
        );
    }
}
