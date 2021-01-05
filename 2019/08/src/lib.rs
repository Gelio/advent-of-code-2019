use image::Image;

pub mod image;

fn find_layer_with_fewest_zeros(img: &Image) -> Option<&Vec<u32>> {
    img.layers.iter().min_by(|img1, img2| {
        let img1_zeros = img1.iter().filter(|d| **d == 0).count();
        let img2_zeros = img2.iter().filter(|d| **d == 0).count();

        img1_zeros.cmp(&img2_zeros)
    })
}

pub fn part_1(img: &Image) -> Option<usize> {
    let layer = find_layer_with_fewest_zeros(img)?;

    let ones_in_layer = layer.iter().filter(|d| **d == 1).count();
    let twos_in_layer = layer.iter().filter(|d| **d == 2).count();

    Some(ones_in_layer * twos_in_layer)
}

pub fn part_2(img: &Image) -> Option<String> {
    let layer = img.get_composed_layer()?;

    Some(layer.to_string())
}
