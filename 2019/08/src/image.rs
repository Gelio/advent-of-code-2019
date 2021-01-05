use std::{error::Error, fmt::Display};

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
}
